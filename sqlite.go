package confgorm

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	DBFile string `env:""`
	db     *gorm.DB
}

func (s *Sqlite) SetDefaults() {
	if s.DBFile == "" {
		// 内存数据库
		s.DBFile = "file::memory:?cache=shared"
	}
}

func (s *Sqlite) Init() {
	if s.db == nil {
		s.SetDefaults()
		s.initial()
	}
}

func (s *Sqlite) initial() {
	db, err := gorm.Open(sqlite.Open(s.DBFile), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connect Sqlite database Failed: %v", err)
	}

	s.db = db
}

func (s *Sqlite) DB() *gorm.DB {
	return s.db
}
