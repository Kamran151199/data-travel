package target

// Mongo is configuration for mongodb.
type Mongo struct {
	Host        string   `json:"host" mapstructure:"host" default:"localhost" validate:"required"`
	Port        int      `json:"port" mapstructure:"port" default:"27017" validate:"required"`
	User        string   `json:"user" mapstructure:"user" default:"guest" validate:"required"`
	Password    string   `json:"password" mapstructure:"password" default:"guest" validate:"required"`
	Database    string   `json:"database" mapstructure:"database" default:"db" validate:"required"`
	Collections []string `json:"collections" mapstructure:"collections" default:"[]" validate:"required"`
}
