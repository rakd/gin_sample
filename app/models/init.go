package models

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
)

var db *gorm.DB

func init() {

	log.Print("models:init")

	dbHost := getenvWithDefault("DBHOST", "localhost")
	dbName := getenvWithDefault("DBNAME", "ginsample")
	dbUser := getenvWithDefault("DBUSER", "root")
	dbPass := getenvWithDefault("DBPASS", "")
	dbPort := getenvWithDefault("DBPORT", "3306")
	if dbPass != "" {
		dbPass = ":" + dbPass
	}

	tryOnlyOnce := getenvWithDefault("TRY_ONLY_ONCE", "")
	skipMigration := getenvWithDefault("SKIP_MIGRATION", "")
	dsn1 := fmt.Sprintf(
		"%s%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	dsn2 := fmt.Sprintf(
		"%s%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		"mysql",
	)

	var dberr error

	loop := 0
	for true {
		db, dberr = gorm.Open("mysql", dsn1)

		log.Printf("dns1:%s", dsn1)
		if dberr == nil {
			log.Printf("db connect OK: %s", dsn1)
			break
		}
		log.Print("DB Connection Error: dsn1=" + dsn1)
		log.Print(dberr.Error())
		if tryOnlyOnce != "" {
			return
		}

		time.Sleep(time.Millisecond * 3000)

		db, dberr = gorm.Open("mysql", dsn2)
		if dberr == nil {
			log.Printf("db connect OK: %s", dsn2)
			log.Printf("create database %s", dbName)
			db.Exec("CREATE DATABASE IF NOT EXISTS `" + dbName + "` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
			// might be error...?
			log.Printf("use dbname")
			db.Exec("use `" + dbName + "` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
			db.Close()
		} else {
			log.Print("DB Connection Error: dsn2=" + dsn2)
			log.Print(dberr.Error())
			loop++
			if loop > 300 {
				//give up to connect... even though no connect db.
				break
			}
		}

	}
	//ResetTables()
	if skipMigration != "" {
		log.Print("skip migration")
	} else {
		log.Printf("before auto migrate.")
		MigrateTables()
		log.Printf("after auto migrate..")
	}
}

// getenvWithDefault ...
func getenvWithDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

// MigrateTables ...
func MigrateTables() {

	log.Print("User")
	db.AutoMigrate(&User{})

}

func generageUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
