package mongo

import (
	"context"
	"fmt"
	errors "github.com/Kamran151199/dbmigrate/pkg/error"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Storage storage implementation.
type Storage struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (s *Storage) GetDB() (*mongo.Database, error) {
	ctx := context.TODO()
	// prepare the connection string
	url := fmt.Sprint("mongodb://", s.User, ":", s.Password, "@", s.Host, ":", s.Port)

	// prepare the default options
	clientOptions := options.Client().ApplyURI(url)

	// create the client connection
	log.Printf("Connecting to mongo server ...\n")
	client, err := mongo.Connect(ctx, clientOptions)

	// if there was an error creating the client, return an error
	if err != nil {
		return nil, errors.CustomError{
			Code:    400,
			Message: fmt.Sprintf("could not get the mongo client %s: ", err),
		}
	}

	// check if the client is alive
	log.Printf("Pinging the mongo server ...\n")
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, errors.CustomError{
			Code:    400,
			Message: fmt.Sprintf("mongo heartbeat failed: %s", err.Error()),
		}
	}

	log.Printf("Connecting specified database ...\n")
	db := client.Database(s.Database)

	if db == nil {
		return nil, errors.CustomError{
			Code:    400,
			Message: fmt.Sprintf("mongo db %s not found: ", s.Database),
		}
	}

	log.Printf("Connected to specified database \n")
	return db, nil
}
