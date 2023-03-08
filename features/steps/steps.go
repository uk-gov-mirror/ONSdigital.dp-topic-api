package steps

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ONSdigital/dp-topic-api/config"
	"github.com/stretchr/testify/assert"

	dpMongoDriver "github.com/ONSdigital/dp-mongodb/v3/mongodb"

	componentModels "github.com/ONSdigital/dp-topic-api/features/models"
	"github.com/ONSdigital/dp-topic-api/models"
	"github.com/cucumber/godog"
	"go.mongodb.org/mongo-driver/bson"
)

func (f *TopicComponent) iHaveTheseTopics(topicsWriteJSON *godog.DocString) error {
	ctx := context.Background()
	var topics []componentModels.TopicWrite
	m := f.MongoClient

	err := json.Unmarshal([]byte(topicsWriteJSON.Content), &topics)
	if err != nil {
		return err
	}

	for _, topicsDoc := range topics {
		if err := f.putTopicInDatabase(ctx, m.Connection.Collection(m.ActualCollectionName(config.TopicsCollection)), topicsDoc); err != nil {
			return err
		}
	}

	return nil
}

func (f *TopicComponent) putTopicInDatabase(ctx context.Context, mongoCollection *dpMongoDriver.Collection, topicDoc componentModels.TopicWrite) error {
	update := bson.M{
		"$set": topicDoc,
		"$setOnInsert": bson.M{
			"last_updated": time.Now(),
		},
	}
	_, err := mongoCollection.UpsertById(ctx, topicDoc.ID, update)
	if err != nil {
		return err
	}
	return nil
}

func (f *TopicComponent) iHaveTheseContents(contentJSON *godog.DocString) error {
	ctx := context.Background()
	var collection []models.ContentResponse
	m := f.MongoClient

	err := json.Unmarshal([]byte(contentJSON.Content), &collection)
	if err != nil {
		return err
	}

	for _, topicsDoc := range collection {
		if err := f.putContentInDatabase(ctx, m.Connection.Collection(m.ActualCollectionName(config.ContentCollection)), topicsDoc); err != nil {
			return err
		}
	}

	return nil
}

func (f *TopicComponent) putContentInDatabase(ctx context.Context, mongoCollection *dpMongoDriver.Collection, contentDoc models.ContentResponse) error {
	update := bson.M{
		"$set": contentDoc,
		"$setOnInsert": bson.M{
			"last_updated": time.Now(),
		},
	}
	_, err := mongoCollection.UpsertById(ctx, contentDoc.ID, update)
	if err != nil {
		return err
	}
	return nil
}

func (f *TopicComponent) privateEndpointsAreEnabled() error {
	f.Config.EnablePrivateEndpoints = true
	return nil
}

func (f *TopicComponent) theDocumentInTheDatabaseForIDShouldBe(documentID string, documentJSON *godog.DocString) error {
	var expectedTopic models.Topic
	currentTime := time.Now()
	startTime := currentTime.Add(-time.Second * 5)

	if err := json.Unmarshal([]byte(documentJSON.Content), &expectedTopic); err != nil {
		return err
	}

	collectionName := f.MongoClient.ActualCollectionName(config.TopicsCollection)
	var actualTopic models.TopicResponse
	if err := f.MongoClient.Connection.Collection(collectionName).FindOne(context.Background(), bson.M{"_id": documentID}, &actualTopic); err != nil {
		return err
	}

	assert.Equal(&f.ErrorFeature, documentID, actualTopic.ID)

	document := actualTopic.Next
	f.ErrorFeature.Log(document)

	// checking last_updated has changed by checking if it was set within the last 5 seconds
	assert.WithinRange(&f.ErrorFeature, *document.LastUpdated, startTime, currentTime)

	// Removing generated timestamps before comparing due to them changing each time tes suite is run
	document.LastUpdated = nil
	expectedTopic.LastUpdated = nil

	assert.Equal(&f.ErrorFeature, expectedTopic, *document)

	return f.ErrorFeature.StepError()
}
