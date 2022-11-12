package database

// https://www.section.io/engineering-education/implement-graphql-api-using-golang-and-mongodb-database/

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"example/graph/model"
)

type DB struct {
	client *mongo.Client
}

func Connect(dbUrl string) *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) FindCardTierScoreById(id string) *model.CardTier {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	cardTiersColl := db.client.Database("SlayTheSpire").Collection("CardsTierList")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res := cardTiersColl.FindOne(ctx, bson.M{"_id": ObjectID})

	cardTier := model.CardTier{ID: id}

	res.Decode(&cardTier)

	return &cardTier
}

func (db *DB) FindCardTierScoreByName(cardName string) *model.CardTier {
	cardTiersColl := db.client.Database("SlayTheSpire").Collection("CardsTierList")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result := cardTiersColl.FindOne(ctx, bson.M{"Card": cardName})

	var cardTier model.CardTier
	result.Decode(&cardTier)

	return &cardTier
}

func (db *DB) CardTiersByClass(class string) []*model.CardTier {
	// cardTiersByClass(class: "Ironclad"){
	// 	id
	// 	card
	// 	overallScore
	//   }

	cardTiersColl := db.client.Database("SlayTheSpire").Collection("CardsTierList")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Find()
	cursor, err := cardTiersColl.Find(ctx, bson.D{{"Class", class}}, opts)
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}

	var cardTiers []*model.CardTier
	for _, result := range results {
		cardTier := model.CardTier{
			ID:      result["_id"].(primitive.ObjectID).Hex(),
			Card:    result["Card"].(string),
			Class:   result["Class"].(string),
			Act1:    int(result["Act1"].(int32)),
			Act2:    int(result["Act2"].(int32)),
			Act3:    int(result["Act3"].(int32)),
			Overall: int(result["Overall"].(int32)),
			Upgrade: int(result["Upgrade"].(int32)),
		}
		cardTiers = append(cardTiers, &cardTier)
	}

	return cardTiers
}
