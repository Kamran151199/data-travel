package connector

// Connector is the interface that defines connector service.
// The connector service is responsible for connecting to a source and destination db.
// This is the initial step in the ETL process.
type Connector interface {
	Connect() error
}
