package connector

import (
	"fmt"
	errors "github.com/Kamran151199/dbmigrate/pkg/error"
	mongodb "github.com/Kamran151199/dbmigrate/pkg/storage/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoConnector struct {
	db     *mongo.Database
	client *mongodb.Storage
}

func (m mongoConnector) Connect() error {
	db, err := m.client.GetDB()
	if err != nil {
		return errors.CustomError{
			Code:    400,
			Message: fmt.Sprintf("could not get the mongo db connection %s: ", err),
		}
	}
	m.db = db
	return nil
}

func NewMongoConnector(dto MongoConnectorDTO) (Connector, error) {
	client := mongodb.Storage{
		Host:     dto.Host,
		Port:     dto.Port,
		User:     dto.User,
		Password: dto.Password,
		Database: dto.Database,
	}

	return &mongoConnector{
		client: &client,
	}, nil
}
