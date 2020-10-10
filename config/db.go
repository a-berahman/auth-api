package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/a-berahman/auth-api/common"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//database
var DB *mongo.Database

//collections
var Users *mongo.Collection
var Sessions *mongo.Collection
var Keys *mongo.Collection

func init() {
	err := godotenv.Load(common.ENV_CONFIG_ADDRESS)
	if err != nil {
		fmt.Println("load enviroment face error : ", err)
	}

	fmt.Println("Connecting to MongoDB...")
	//get a mongo session
	url := ""
	dbName := ""
	if os.Getenv(common.MODE) == common.DEV_MODE {

		url = os.Getenv(common.DEV_DB_URL)
		dbName = os.Getenv(common.DEV_DB_NAME)

	} else if os.Getenv(common.MODE) == common.PRD_MODE {
		url = os.Getenv(common.PRD_DB_URL)
		dbName = os.Getenv(common.PRD_DB_NAME)
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// if err = s.Ping(); err != nil {
	// 	panic(err)
	// }

	Users = client.Database(dbName).Collection("users")
	Sessions = client.Database(dbName).Collection("sessions")
	Keys = client.Database(dbName).Collection("keys")

	fmt.Println("connecting to mongo database.")

}
