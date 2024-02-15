package config

import (
	"gorm.io/gorm"
)

var (
	logger   *Logger
	db       *gorm.DB
	messages map[string]string
)

func InitializeConfigs() error {
	var err error

	db, err = InitializeMysqlDb()
	if err != nil {
		logger.ErrF("Failed to connect with database: ", err)
	}

	messages, err = ReadMessagesPropertiesFile()
	if err != nil {
		logger.ErrF("Failed retrieve error messages: ", err)
	}

	return err

}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}

func GetMysqlDb() *gorm.DB {
	return db
}

func FindErrorMessage(errorMessage string) string {
	return messages[errorMessage]
}
