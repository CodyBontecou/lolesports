package service

import "time"

type LiveEvents struct {
	Data struct {
		Schedule struct {
			Events []*Event `json:"events" bson:"events,omitempty"`
		} `json:"schedule" bson:"schedule,omitempty"`
	} `json:"data" bson:"data,omitempty"`
}

type Event struct {
	Id        string    `json:"id" bson:"id,omitempty"`
	Type      string    `json:"type" bson:"type,omitempty"`
	State     string    `json:"state" bson:"state,omitempty"`
	StartTime time.Time `json:"startTime" bson:"startTime,omitempty"`
	BlockName string    `json:"blockName" bson:"blockName,omitempty"`
	Match     *Match    `json:"match,omitempty" bson:"match,omitempty"`
	League    *League   `json:"league,omitempty" bson:"league,omitempty"`
}

type League struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Slug     string `json:"slug,omitempty" bson:"slug,omitempty"`
	Id       string `json:"id,omitempty" bson:"id,omitempty"`
	Image    string `json:"image,omitempty" bson:"image,omitempty"`
	Priority int    `json:"priority,omitempty" bson:"priority,omitempty"`
}

type Match struct {
	Teams    []*Team `json:"teams,omitempty" bson:"teams,omitempty"`
	Id       string  `json:"id,omitempty" bson:"id,omitempty"`
	Strategy struct {
	} `json:"strategy,omitempty" bson:"strategy,omitempty"`
	Games []*Game `json:"games,omitempty,omitempty" bson:"games,omitempty"`
}

type Team struct {
	Id     string `json:"id" bson:"id,omitempty"`
	Name   string `json:"name" bson:"name,omitempty"`
	Slug   string `json:"slug" bson:"slug,omitempty"`
	Code   string `json:"code" bson:"code,omitempty"`
	Image  string `json:"image" bson:"image,omitempty"`
	Result struct {
		Outcome  interface{} `json:"outcome" bson:"outcome,omitempty"`
		GameWins int         `json:"gameWins" bson:"gameWins"`
	} `json:"result" bson:"result,omitempty"`
	Record struct {
		Wins   int `json:"wins" bson:"wins"`
		Losses int `json:"losses" bson:"losses"`
	} `json:"record" bson:"record,omitempty"`
}

type Game struct {
	Number int    `json:"number" bson:"number,omitempty"`
	Id     string `json:"id" bson:"id,omitempty"`
	State  string `json:"state" bson:"state,omitempty"`
	Teams  []struct {
		Id   string `json:"id" bson:"id,omitempty"`
		Side string `json:"side" bson:"side,omitempty"`
	} `json:"teams" bson:"teams,omitempty"`
	Vods []interface{} `json:"vods" bson:"vods,omitempty"`
}

type Frames struct {
	Frames []*Frame `json:"frames" bson:"frames,omitempty"`
}

type GameMetadata struct {
	PatchVersion     string       `json:"patchVersion" bson:"patchVersion,omitempty"`
	BlueTeamMetadata TeamMetadata `json:"blueTeamMetadata" bson:"blueTeamMetadata,omitempty"`
	RedTeamMetadata  TeamMetadata `json:"redTeamMetadata" bson:"redTeamMetadata,omitempty"`
}

type TeamMetadata struct {
	EsportsTeamId       string                 `json:"esportsTeamId" bson:"esportsTeamId,omitempty"`
	ParticipantMetadata []*ParticipantMetadata `json:"participantMetadata" bson:"participantMetadata,omitempty"`
}

type ParticipantMetadata struct {
	ParticipantId   int    `json:"participantId" bson:"participantId,omitempty"`
	EsportsPlayerId string `json:"esportsPlayerId" bson:"esportsPlayerId,omitempty"`
	SummonerName    string `json:"summonerName" bson:"summonerName,omitempty"`
	ChampionId      string `json:"championId" bson:"championId,omitempty"`
	Role            string `json:"role" bson:"role,omitempty"`
}

type Frame struct {
	Rfc460Timestamp time.Time     `json:"rfc460Timestamp" bson:"rfc460Timestamp,omitempty"`
	Participants    []Participant `json:"participants" bson:"participants,omitempty"`
}

type Participant struct {
	ParticipantId       int           `json:"participantId" bson:"participantId"`
	Level               int           `json:"level" bson:"level"`
	Kills               int           `json:"kills" bson:"kills"`
	Deaths              int           `json:"deaths" bson:"deaths"`
	Assists             int           `json:"assists" bson:"assists"`
	TotalGoldEarned     int           `json:"totalGoldEarned" bson:"totalGoldEarned"`
	CreepScore          int           `json:"creepScore" bson:"creepScore"`
	KillParticipation   float64       `json:"killParticipation" bson:"killParticipation"`
	ChampionDamageShare float64       `json:"championDamageShare" bson:"championDamageShare"`
	WardsPlaced         int           `json:"wardsPlaced" bson:"wardsPlaced"`
	WardsDestroyed      int           `json:"wardsDestroyed" bson:"wardsDestroyed"`
	AttackDamage        int           `json:"attackDamage" bson:"attackDamage"`
	AbilityPower        int           `json:"abilityPower" bson:"abilityPower"`
	CriticalChance      float64       `json:"criticalChance" bson:"criticalChance"`
	AttackSpeed         int           `json:"attackSpeed" bson:"attackSpeed"`
	LifeSteal           int           `json:"lifeSteal" bson:"lifeSteal"`
	Armor               int           `json:"armor" bson:"armor"`
	MagicResistance     int           `json:"magicResistance" bson:"magicResistance"`
	Tenacity            float64       `json:"tenacity" bson:"tenacity,omitempty"`
	Items               []interface{} `json:"items" bson:"items,omitempty"`
	PerkMetadata        *PerkMetadata `json:"perkMetadata,omitempty" bson:"-"`
	Abilities           []interface{} `json:"abilities" bson:"abilities,omitempty"`
}

type PerkMetadata struct {
	StyleId    int   `json:"styleId" bson:"styleId,omitempty"`
	SubStyleId int   `json:"subStyleId" bson:"subStyleId,omitempty"`
	Perks      []int `json:"perks" bson:"perks,omitempty"`
}

type Window struct {
	EsportsGameId  string        `json:"esportsGameId" bson:"esportsGameId,omitempty"`
	EsportsMatchId string        `json:"esportsMatchId" bson:"esportsMatchId,omitempty"`
	GameMetadata   *GameMetadata `json:"gameMetadata" bson:"gameMetadata,omitempty"`
	Frames         []WindowFrame `json:"frames" bson:"frames,omitempty"`
}

type WindowFrame struct {
	Rfc460Timestamp time.Time       `json:"rfc460Timestamp" bson:"rfc460Timestamp,omitempty"`
	GameState       string          `json:"gameState" bson:"gameState,omitempty"`
	BlueTeam        WindowFrameTeam `json:"blueTeam" bson:"blueTeam,omitempty"`
	RedTeam         WindowFrameTeam `json:"redTeam" bson:"redTeam,omitempty"`
}

type WindowFrameTeam struct {
	TotalGold    int                       `json:"totalGold" bson:"totalGold"`
	Inhibitors   int                       `json:"inhibitors" bson:"inhibitors"`
	Towers       int                       `json:"towers" bson:"towers"`
	Barons       int                       `json:"barons" bson:"barons"`
	TotalKills   int                       `json:"totalKills" bson:"totalKills"`
	Dragons      []interface{}             `json:"dragons" bson:"dragons"`
	Participants []*WindowFrameParticipant `json:"participants" bson:"participants,omitempty"`
}

type WindowFrameParticipant struct {
	ParticipantId int `json:"participantId" bson:"participantId"`
	TotalGold     int `json:"totalGold" bson:"totalGold"`
	Level         int `json:"level" bson:"level"`
	Kills         int `json:"kills" bson:"kills"`
	Deaths        int `json:"deaths" bson:"deaths"`
	Assists       int `json:"assists" bson:"assists"`
	CreepScore    int `json:"creepScore" bson:"creepScore"`
	CurrentHealth int `json:"currentHealth" bson:"currentHealth"`
	MaxHealth     int `json:"maxHealth" bson:"maxHealth"`
}
