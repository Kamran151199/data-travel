package connector

import (
	"fmt"
	errors "github.com/Kamran151199/dbmigrate/pkg/error"
	"github.com/Kamran151199/dbmigrate/pkg/storage/postgresql"
	"github.com/jmoiron/sqlx"
	"log"
)

type postgresqlConnector struct {
	db     *sqlx.DB
	client *postgresql.Storage
}

func NewPostgresqlConnector(dto PostgresqlConnectorDTO) (Connector, error) {
	client := postgresql.Storage{
		Host:     dto.Host,
		Port:     dto.Port,
		Database: dto.Database,
		User:     dto.User,
		Password: dto.Password,
	}

	return &postgresqlConnector{
		client: &client,
	}, nil
}

func (c *postgresqlConnector) Connect() error {
	// get the db connection
	log.Println("Connecting to postgresql ...")
	db, err := c.client.GetDB()

	// if there was an error creating the connection pool, return an error
	if err != nil {
		return errors.CustomError{
			Code:    400,
			Message: fmt.Sprintf("could not get the postgresql connection pool %s: ", err),
		}
	}

	log.Printf("Connected to postgresql \n")

	// ping the db to check if it is alive
	log.Printf("Pinging the postgresql db ...\n")
	err = c.db.Ping()

	// if there was an error pinging the db, return an error
	if err != nil {
		return errors.CustomError{
			Code:    400,
			Message: fmt.Sprintf("postgresql heartbeat failed: %s", err.Error()),
		}
	}
	c.db = db
	log.Println("Postgresql db is up and running")

	return nil
}
