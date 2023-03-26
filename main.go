package main

import (
	"database/sql"
	"fmt"
	"os"
	"project-trial/controllers"
	"project-trial/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	//config env
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file env.")
	} else {
		fmt.Println("success read file env.")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))
	fmt.Println(psqlInfo)
	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}
	database.DbMigrate(DB)
	defer DB.Close()
	//Router
	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("persons/:id", controllers.UpdatePerson)
	router.DELETE("persons/:id", controllers.DeletePerson)
	router.Run(":" + os.Getenv("PORT"))

}
