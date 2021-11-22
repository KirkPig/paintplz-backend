package mongo_repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//

// Config config
type Config struct {
	Host string
	Db   string
}

type MongoClient interface {
	ping(ctx context.Context, rp *readpref.ReadPref) error
	database() *mongo.Database
}

type mongoClient struct {
	client *mongo.Client
	config Config
}

func (c *mongoClient) ping(ctx context.Context, rp *readpref.ReadPref) error {
	return c.client.Ping(ctx, nil)
}

func (c *mongoClient) database() *mongo.Database {
	return c.client.Database(c.config.Db)
}

func NewMongoClient(c Config) MongoClient {
	clientOptions := options.Client().ApplyURI(c.Host)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// return client.Database(c.Db)
	return &mongoClient{
		client: client,
		config: c,
	}
}

var (
	ArtworkCollection       *mongo.Collection
	ArtworkArtTagCollection *mongo.Collection
)

func GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*30)
}

func Init() {
	config := Config{Host: "mongodb://root:123456@localhost", Db: "paintplzIO"}
	client := NewMongoClient(config)
	err := client.ping(context.TODO(), nil)
	if err != nil {
		log.Panic(err)
	}
	ArtworkCollection = client.database().Collection("artwork")
	ArtworkArtTagCollection = client.database().Collection("artwork_arttag")
	log.Println("Mongodb connected!! 🎉")
}
