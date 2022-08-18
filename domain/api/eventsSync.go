package api

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"lolesports/domain/db"
	"lolesports/domain/db/events"
	"lolesports/domain/db/frames"
	api "lolesports/domain/service"
	"runtime"
	"sync"
	"time"
)

const roundValue = 10 * time.Second

type EventSyncer struct {
	Mu         sync.Mutex
	LiveEvents map[string]*events.AllEventData
}

func NewEventSyncer(service *db.UseCase) *EventSyncer {
	es := &EventSyncer{
		LiveEvents: make(map[string]*events.AllEventData),
	}

	liveEvents, err := service.Repository.EventsRepo.Find(bson.M{"event.state": "inProgress"}, nil)
	if err != nil {
		panic(err.Error())
	}
	for _, event := range liveEvents {
		event := event
		es.LiveEvents[event.Event.Id] = event
	}

	return es
}

func (es *EventSyncer) SyncLiveEvents(service *db.UseCase) {
	for {
		liveEvents, err := api.GetLiveEvents()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		es.Mu.Lock()
		for key, value := range es.LiveEvents {
			value.Event.State = "ended"
			es.LiveEvents[key] = value
		}

		for _, value := range liveEvents.Data.Schedule.Events {
			if value.Type == "show" {
				continue
			}
			if eventData, exists := es.LiveEvents[value.Id]; !exists {

				newEvent := &events.AllEventData{
					Event: value,
					Games: make(map[string]*events.GameData),
				}

				newEvent, err = service.Repository.EventsRepo.Add(context.Background(), *newEvent)
				es.LiveEvents[newEvent.Event.Id] = newEvent
			} else {
				eventData.Event = value
				es.LiveEvents[value.Id] = eventData
			}
		}

		for key, value := range es.LiveEvents {
			if value.Event.State == "ended" {
				service.Repository.EventsRepo.Add(context.Background(), *value)
				delete(es.LiveEvents, key)
			}
		}
		es.Mu.Unlock()
		fmt.Println("sync live events")
		time.Sleep(10 * time.Second)
	}
}

func (es *EventSyncer) GetLiveEventsData(service *db.UseCase) {
	var wg sync.WaitGroup
	for {
		es.Mu.Lock()
		for _, event := range es.LiveEvents {
			event := event
			wg.Add(1)
			go func() {
				defer wg.Done()
				es.getLiveEventData(service, event)
			}()
		}
		wg.Wait()
		es.Mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func (es *EventSyncer) getLiveEventData(service *db.UseCase, event *events.AllEventData) {
	eventFilter := bson.D{{"event.id", event.Event.Id}}
	if event.Event.Match == nil || event.Event.Match.Games == nil {
		return
	}

	for _, game := range event.Event.Match.Games {
		if event.Games == nil {
			event.Games = make(map[string]*events.GameData)
		}

		gameData, exists := event.Games[game.Id]
		if !exists {
			gameData = &events.GameData{
				PerksMetadata: map[int]*api.PerkMetadata{},
			}
			event.Games[game.Id] = gameData
		}

		if gameData.Fetched {
			continue
		}

		if gameData.LastTimeChecked != nil {
			aux := gameData.LastTimeChecked.Add(10 * time.Second)
			gameData.LastTimeChecked = &aux
		}

		// WINDOW
		window, err := api.GetWindow(game.Id, gameData.LastTimeChecked)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if gameData.GameMetadata == nil {
			for index, participant := range window.GameMetadata.BlueTeamMetadata.ParticipantMetadata {
				champ, _ := service.Repository.ChampsRepo.FindOne(bson.D{{"_id", participant.ChampionId}})
				participant.ChampionId = champ.Key
				window.GameMetadata.BlueTeamMetadata.ParticipantMetadata[index] = participant
			}

			for index, participant := range window.GameMetadata.RedTeamMetadata.ParticipantMetadata {
				champ, _ := service.Repository.ChampsRepo.FindOne(bson.D{{"_id", participant.ChampionId}})
				participant.ChampionId = champ.Key
				window.GameMetadata.RedTeamMetadata.ParticipantMetadata[index] = participant
			}

			gameData.GameMetadata = window.GameMetadata

			set := bson.D{{"$set", bson.M{fmt.Sprintf("games.%s.game_metadata", game.Id): gameData.GameMetadata}}}
			service.Repository.EventsRepo.UpdateMany(eventFilter, set)
		}

		for _, frame := range window.Frames {
			if frame.GameState == "finished" {
				gameData.Fetched = true
				set := bson.D{{"$set", bson.M{fmt.Sprintf("games.%s.fetched", game.Id): true}}}
				service.Repository.EventsRepo.UpdateMany(eventFilter, set)
			}
			if gameData.LastTimeChecked == nil {
				var roundedTime = frame.Rfc460Timestamp.Round(roundValue)
				gameData.LastTimeChecked = &roundedTime
			} else if frame.Rfc460Timestamp.Sub(*gameData.LastTimeChecked) < 0 {
				continue
			}

			// check if game started
			if frame.BlueTeam.TotalGold != 0 {
				if gameData.GameStartTime == nil {
					gameData.GameStartTime = &frame.Rfc460Timestamp
					set := bson.D{{"$set", bson.M{fmt.Sprintf("games.%s.game_start_time", game.Id): gameData.GameStartTime}}}
					service.Repository.EventsRepo.UpdateMany(eventFilter, set)
				}

				if gameData.PausedAt != nil && frame.GameState == "in_game" && frame.Rfc460Timestamp.Sub(*gameData.PausedAt) > 0 {
					gameData.PausedDuration = frame.Rfc460Timestamp.Sub(*gameData.PausedAt)
					gameData.PausedAt = nil

					set := bson.D{{"$set", bson.D{
						{fmt.Sprintf("games.%s.paused_at", game.Id), nil},
						{fmt.Sprintf("games.%s.paused_duration", game.Id), gameData.PausedDuration},
					}}}
					service.Repository.EventsRepo.UpdateMany(eventFilter, set)
				}

				mapPos := int((frame.Rfc460Timestamp.Sub(*gameData.GameStartTime) - gameData.PausedDuration).Seconds())
				if mapPos <= gameData.LastWindowFrame {
					continue
				}

				gameData.LastWindowFrame = mapPos
				wFrame := frames.AllFrameData{
					EventID:     event.Event.Id,
					GameID:      game.Id,
					Second:      mapPos,
					WindowFrame: &frame,
				}
				service.Repository.FramesRepo.Add(context.Background(), wFrame)
				if err != nil {
					return
				}

				set := bson.D{{"$set", bson.M{fmt.Sprintf("games.%s.last_window_frame", game.Id): gameData.LastWindowFrame}}}
				service.Repository.EventsRepo.UpdateMany(eventFilter, set)

				if frame.GameState == "paused" {
					gameData.PausedAt = &frame.Rfc460Timestamp

					set := bson.D{{"$set", bson.M{fmt.Sprintf("games.%s.paused_at", game.Id): gameData.PausedAt}}}
					service.Repository.EventsRepo.UpdateMany(eventFilter, set)
				}

			}
		}

		// DETAIL
		detailFrames, err := api.GetDetails(game.Id, gameData.LastTimeChecked)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if len(gameData.PerksMetadata) == 0 && len(detailFrames) > 0 {
			if gameData.PerksMetadata == nil {
				gameData.PerksMetadata = make(map[int]*api.PerkMetadata)
			}
			for _, participant := range detailFrames[0].Participants {
				gameData.PerksMetadata[participant.ParticipantId] = participant.PerkMetadata
			}
			set := bson.D{{"$set", bson.D{
				{fmt.Sprintf("games.%s.perks_metadata", game.Id), gameData.PerksMetadata},
			}}}
			service.Repository.EventsRepo.UpdateMany(eventFilter, set)
		}

		for _, frame := range detailFrames {
			if frame.Rfc460Timestamp.Sub(*gameData.LastTimeChecked) < 0 {
				continue
			}
			// check if game started
			if frame.Participants[0].TotalGoldEarned != 0 {
				if gameData.GameStartTime == nil {
					gameData.GameStartTime = &frame.Rfc460Timestamp

					set := bson.D{{"$set", bson.M{fmt.Sprintf("games.%s.game_start_time", game.Id): gameData.GameStartTime}}}
					service.Repository.EventsRepo.UpdateMany(eventFilter, set)
				}
				mapPos := int((frame.Rfc460Timestamp.Sub(*gameData.GameStartTime) - gameData.PausedDuration).Seconds())

				if mapPos <= gameData.LastFrame {
					continue
				}

				gameData.LastFrame = mapPos
				wFrame := frames.AllFrameData{
					EventID: event.Event.Id,
					GameID:  game.Id,
					Second:  mapPos,
					Frame:   frame,
				}
				_, err = service.Repository.FramesRepo.Add(context.Background(), wFrame)
				if err != nil {
					return
				}

				set := bson.D{{"$set", bson.M{fmt.Sprintf("games.%s.last_frame", game.Id): gameData.LastFrame}}}
				service.Repository.EventsRepo.UpdateMany(eventFilter, set)
			}
		}

		if gameData.LastTimeChecked != nil {
			set := bson.D{{"$set", bson.M{fmt.Sprintf("games.%s.last_time_checked", game.Id): gameData.LastTimeChecked}}}
			service.Repository.EventsRepo.UpdateMany(eventFilter, set)
		}
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
