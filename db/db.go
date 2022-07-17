package db

import (
	"sovereign/configs"
	"sovereign/models"
	"sovereign/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dbUri := configs.DBConnStr()
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbUri,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	utils.HandleErr(err, "Could not connect to DB at URI")

	refreshSchema(DB)
	return DB
}

func refreshSchema(db *gorm.DB) {
	// this drops all of this projects tables
	// re-creates them with the defined schema
	// and seeds some data
	var tables []interface{}
	tables = append(tables, &models.Flag{}, &models.Audience{}, &models.Attribute{}, &models.Condition{})

	// drop all relevant tables
	db.Migrator().DropTable(tables...)
	db.Migrator().DropTable("flag_audiences")

	// create all relevant tables
	db.AutoMigrate(tables...)
	// seed some data
	seedDB(db)
}