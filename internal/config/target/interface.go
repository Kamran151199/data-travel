package target

// Target is a configuration of either source or sink database.
type Target interface {
	Mongo | Postgresql
}
