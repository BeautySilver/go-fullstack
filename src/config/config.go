package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var(
	PORT =0
	SECRETKEY []byte
	DBDRIVER=""
	DBURL=""
)

func Load(){
	var err error
	err = godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err!=nil{
		log.Println(err)
		PORT = 9000
	}
	DBHOST := os.Getenv("DATABASE_HOST")
	if DBHOST ==""{
		DBHOST="127.0.0.1"
	}

    DBDRIVER= os.Getenv("DB_DRIVER")
	DBURL=fmt.Sprintf("%s:%s@tcp(127.0.01:3306)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		 os.Getenv("DB_NAME"))
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
}