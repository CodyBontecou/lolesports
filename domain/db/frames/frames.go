package frames

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lolesports/utils"
	"strings"
	"time"
)

const (
	mongoQueryTimeout          = 10 * time.Second
	frameCollectionName string = "frames"
)

func (repository *scenesRepository) Add(ctx context.Context, frame AllFrameData) (*AllFrameData, error) {
	var cancel context.CancelFunc
	if ctx == nil {
		ctx, cancel = context.WithTimeout(context.Background(), mongoQueryTimeout)
		defer cancel()
	}

	logCTX := utils.BaseLogCtx()

	options := options.FindOneAndUpdate().SetUpsert(true)

	response := repository.db.Collection(frameCollectionName).FindOneAndUpdate(
		ctx,
		bson.M{"event_id": frame.EventID, "second": frame.Second, "game_id": frame.GameID},
		bson.M{"$set": frame},
		options)

	if response.Err() != nil && response.Err().Error() != "mongo: no documents in result" {
		utils.WithFields(logCTX, log.Fields{
			"context": "frame_repository_add",
			"error":   strings.ReplaceAll(response.Err().Error(), " ", "_"),
		})
		logCTX.Level = log.ErrorLevel
		repository.logger.HandleLog(logCTX)
		return nil, response.Err()
	}

	response.Decode(&frame)

	utils.WithFields(logCTX, log.Fields{
		"context": "frame_repository_add",
		"detail":  fmt.Sprintf("Saved_frame: frame_%s, second_%d", frame.EventID, frame.Second),
	})
	repository.logger.HandleLog(logCTX)

	return &frame, nil
}

func (repository *scenesRepository) UpdateMany(filter bson.D, set bson.D) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	logCTX := utils.BaseLogCtx()
	_, err := repository.db.Collection(frameCollectionName).UpdateMany(
		ctx,
		filter,
		set,
	)

	if err != nil {
		utils.WithFields(logCTX, log.Fields{
			"context": "frame_update_many",
			"error":   strings.ReplaceAll(err.Error(), " ", "_"),
		})
		logCTX.Level = log.ErrorLevel
		repository.logger.HandleLog(logCTX)
		return err
	}
	return nil
}

func (repository *scenesRepository) Find(filter bson.M, findOption *options.FindOptions) ([]*AllFrameData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	logCTX := utils.BaseLogCtx()
	response, err := repository.db.Collection(frameCollectionName).Find(
		ctx,
		filter,
		findOption,
	)

	if err != nil {
		utils.WithFields(logCTX, log.Fields{
			"context": "frames_find",
			"error":   strings.ReplaceAll(err.Error(), " ", "_"),
		})
		logCTX.Level = log.ErrorLevel
		repository.logger.HandleLog(logCTX)
		return nil, err
	}
	defer response.Close(ctx)

	var frames []*AllFrameData
	err = response.All(ctx, &frames)
	if err != nil {
		utils.WithFields(logCTX, log.Fields{
			"context": "frames_find",
			"error":   strings.ReplaceAll(err.Error(), " ", "_"),
		})
		logCTX.Level = log.ErrorLevel
		repository.logger.HandleLog(logCTX)
		return nil, err
	}

	return frames, nil
}

func (repository *scenesRepository) FindOne(filter bson.M) (*AllFrameData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	logCTX := utils.BaseLogCtx()

	response := repository.db.Collection(frameCollectionName).FindOne(
		ctx,
		filter,
	)

	if response.Err() != nil {
		return nil, response.Err()
	}

	var frame *AllFrameData
	err := response.Decode(&frame)
	if err != nil {
		utils.WithFields(logCTX, log.Fields{
			"context": "frame_parse",
			"error":   strings.ReplaceAll(response.Err().Error(), " ", "_"),
		})
		logCTX.Level = log.ErrorLevel
		repository.logger.HandleLog(logCTX)
		return nil, response.Err()
	}

	return frame, nil
}
