package migration

func NewDatabase(name string) *Database {
	return &Database{
		name:   name,
		tables: make([]interface{}, 0),
	}
}

type Database struct {
	name   string
	tables []interface{}
}

func (db *Database) AddTable(table interface{}) {
	// todo: checking table
	db.tables = append(db.tables, table)
}

func (db *Database) Tables() []interface{} {
	return db.tables
}

func (db *Database) Name() string {
	return db.name
}
