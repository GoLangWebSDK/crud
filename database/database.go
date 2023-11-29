package database

type Database struct {
	Adapter Adapter
}

func New(adapter Adapter) *Database {
	return &Database{
		Adapter: adapter,
	}
}