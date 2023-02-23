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

	if err := json.Unmarshal([]byte(documentJSON.Content), &expectedTopic); err != nil {
		return err
	}

	collectionName := f.MongoClient.ActualCollectionName(config.TopicsCollection)
	var link models.TopicResponse
	if err := f.MongoClient.Connection.Collection(collectionName).FindOne(context.Background(), bson.M{"_id": documentID}, &link); err != nil {
		return err
	}

	assert.Equal(&f.ErrorFeature, documentID, link.ID)

	document := link.Next

	assert.Equal(&f.ErrorFeature, expectedTopic, *document)

	return f.ErrorFeature.StepError()
}
