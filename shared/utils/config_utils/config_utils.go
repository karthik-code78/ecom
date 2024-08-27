package config_utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func LoadEnv() string {
	wd, err := os.Getwd()
	if strings.Contains(wd, "cmd") {
		wd = filepath.Join(wd, "../")
	}
	log.Print(wd)
	if err != nil {
		log.Fatal("Unable to get working directory")
	}
	err = godotenv.Load(path.Join(wd, "/.env"))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return wd
}
