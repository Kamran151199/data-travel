package target

// Postgresql is configuration for PostgreSQL database.
type Postgresql struct {
	Host     string   `json:"host" mapstructure:"host" default:"localhost" validate:"required"`
	Port     int      `json:"port" mapstructure:"port" default:"5432" validate:"required"`
	User     string   `json:"user" mapstructure:"user" default:"postgres" validate:"required"`
	Password string   `json:"password" mapstructure:"password" default:"postgres" validate:"required"`
	Database string   `json:"database" mapstructure:"database" default:"postgres" validate:"required"`
	Tables   []string `json:"tables" mapstructure:"tables" validate:"required"`
}
