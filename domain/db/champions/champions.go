package champions

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/apex/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"lolesports/utils"
)

const (
	mongoQueryTimeout              = 10 * time.Second
	championsCollectionName string = "champions"
)

func (repository *championsRepository) Add(champion *Champion) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	logCTX := utils.BaseLogCtx()

	err := repository.db.Collection(championsCollectionName).FindOneAndUpdate(
		ctx,
		bson.M{"_id": champion.ID},
		bson.M{"$set": champion},
		options.FindOneAndUpdate().SetUpsert(true),
	).Err()

	if err != nil && err.Error() != "mongo: no documents in result" {
		utils.WithFields(logCTX, log.Fields{
			"context": "champion_repository_add",
			"error":   strings.ReplaceAll(err.Error(), " ", "_"),
		})
		logCTX.Level = log.ErrorLevel
		repository.logger.HandleLog(logCTX)
		return err
	}

	utils.WithFields(logCTX, log.Fields{
		"context": "champion_repository_add",
		"detail":  fmt.Sprintf("Saved_champion:%s", champion.ID),
	})
	repository.logger.HandleLog(logCTX)
	return nil
}

func (repository *championsRepository) FindOne(filter bson.D) (*Champion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	logCTX := utils.BaseLogCtx()

	response := repository.db.Collection(championsCollectionName).FindOne(
		ctx,
		filter,
	)

	if response.Err() != nil {
		fmt.Println(filter)
		utils.WithFields(logCTX, log.Fields{
			"context": "champion_find",
			"error":   strings.ReplaceAll(response.Err().Error(), " ", "_"),
		})
		logCTX.Level = log.ErrorLevel
		repository.logger.HandleLog(logCTX)
		return nil, response.Err()
	}

	var champion Champion
	response.Decode(&champion)

	return &champion, nil
}

func (repository *championsRepository) Find(filter bson.M, findOption *options.FindOptions) ([]Champion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoQueryTimeout)
	defer cancel()

	logCTX := utils.BaseLogCtx()

	response, err := repository.db.Collection(championsCollectionName).Find(
		ctx,
		filter,
		findOption,
	)

	if err != nil {
		utils.WithFields(logCTX, log.Fields{
			"context": "account_find",
			"error":   strings.ReplaceAll(err.Error(), " ", "_"),
		})
		logCTX.Level = log.ErrorLevel
		repository.logger.HandleLog(logCTX)
		return nil, err
	}
	defer response.Close(ctx)

	var champions []Champion
	err = response.All(ctx, &champions)
	if err != nil {
		fmt.Println(filter)
		utils.WithFields(logCTX, log.Fields{
			"context": "champion_find",
			"error":   strings.ReplaceAll(err.Error(), " ", "_"),
		})
		logCTX.Level = log.ErrorLevel
		repository.logger.HandleLog(logCTX)
		return nil, err
	}

	return champions, nil
}
