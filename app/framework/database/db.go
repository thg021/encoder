package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"github.com/thg021/encoder/domain"
)

type Database struct {
	Db            *gorm.DB
	Dns           string
	DnsTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "dev"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DnsTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test db error: %s", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {

	var err error
	if d.Env != "dev" {
		d.Db, err = gorm.Open(d.DbType, d.Dns)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DnsTest)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
		d.Db.Raw("PRAGMA foreign_keys=ON")
		d.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos (id)", "CASCADE", "CASCADE")
	}

	return d.Db, nil
}
