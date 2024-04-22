package config

import (
	"fmt"
	"log"

	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbMap *gorm.DB

func DB() *gorm.DB {
	if dbMap == nil {
		dbMap = connectDatabase()
	}
	return dbMap
}

func connectDatabase() *gorm.DB {
	envHasLog := GetEnvVariable("DB_HAS_LOG", "false")
	shouldLog, _ := strconv.ParseBool(envHasLog)
	logLevel := logger.Silent
	if shouldLog {
		logLevel = logger.Info
	}
	connString := getConnectionString()
	var err error
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		logrus.Errorf("Error in connecting to db: %s", err.Error())
		panic("Failed to connect to the database!")
	}

	log.Println("Database connection established")
	return db
}

func getConnectionString() string {
	var connString strings.Builder
	connString.WriteString(GetEnvVariable("DB_USER", "user"))
	connString.WriteString(":")
	connString.WriteString(GetEnvVariable("DB_PASSWORD", "Password!3306"))
	connString.WriteString("@tcp(")
	connString.WriteString(GetEnvVariable("DB_HOST", "host.docker.internal"))
	connString.WriteString(":")
	connString.WriteString(GetEnvVariable("DB_PORT", "3306"))
	connString.WriteString(")/")
	connString.WriteString(GetEnvVariable("DB_NAME", "atmail"))
	connString.WriteString("?charset=utf8")
	connString.WriteString("&parseTime=True")
	connString.WriteString("&loc=Local")
	fmt.Println(connString.String())
	return connString.String()
}
