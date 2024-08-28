package migration

import (
	"github.com/karthik-code78/ecom/shared/configure"
	"github.com/karthik-code78/ecom/shared/logging"
)

func Migrate(models ...interface{}) {
	db, err := configure.ConnectAndReturnDB()
	if err != nil {
		logging.Log.Fatal("Failed to connect to the Database", err)
	}
	logging.Log.Info("DB in migrate is: ", db)
	err = db.AutoMigrate(models...)
	if err != nil {
		logging.Log.Error("Failed to migrate DB", err)
	}
}
