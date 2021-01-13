package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/devmeireles/gnosi-api/app/routes"
	"github.com/devmeireles/gnosi-api/app/utils"
	"github.com/joho/godotenv"
)

var server = routes.Server{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		fmt.Println(err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		log.Fatal("$DB_DRIVER must be set")
	}

	dbUser := os.Getenv("DB_USER_TEST")
	if dbUser == "" {
		log.Fatal("$DB_USER_TEST must be set")
	}

	dbPassword := os.Getenv("DB_PASSWORD_TEST")
	if dbPassword == "" {
		log.Fatal("$DB_PASSWORD_TEST must be set")
	}

	dbPort := os.Getenv("DB_PORT_TEST")
	if dbPort == "" {
		log.Fatal("$DB_PORT_TEST must be set")
	}

	dbHost := os.Getenv("DB_HOST_TEST")
	if dbHost == "" {
		log.Fatal("$DB_HOST_TEST must be set")
	}

	dbName := os.Getenv("DB_NAME_TEST")
	if dbName == "" {
		log.Fatal("$DB_NAME_TEST must be set")
	}

	utils.InitDatabase(
		dbDriver,
		dbUser,
		dbPassword,
		dbPort,
		dbHost,
		dbName,
	)

	server.SetupRoutes()

	os.Exit(m.Run())
}
