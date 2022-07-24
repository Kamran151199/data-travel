package domain

type Mongo struct {
	Host     string `json:"host,omitempty" default:"localhost"`
	Port     string `json:"port,omitempty" default:"27017"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
}

type Postgres struct {
	Host     string `json:"host,omitempty" default:"localhost"`
	Port     string `json:"port,omitempty" default:"5432"`
	User     string `json:"user,omitempty" default:"postgres"`
	Password string `json:"password,omitempty" default:"password"`
	Database string `json:"database,omitempty" default:"postgres"`
	MaxConn  int    `json:"max_conn,omitempty" default:"10"`
}

type Connection interface {
	Mongo | Postgres
}

type ETL[O Connection, D Connection] struct {
	Origin            O      `json:"origin"`
	Destination       D      `json:"destination"`
	OriginTarget      string `json:"origin_target"`
	DestinationTarget string `json:"destination_target"`
}
