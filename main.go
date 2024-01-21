package main

import (
	"dbo-test/handlers"
	"dbo-test/helpers/env"
	"dbo-test/models"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	// set up db connection
	cs := env.Get().ConnectionString()
	db, err := gorm.Open(mysql.Open(cs), &gorm.Config{})

	//check DB Connection
	if err != nil {
		log.Fatal("Connection To DB Failed")
		return
	}
	log.Println("DB Connected Successfully")

	//migrate db structured
	err = models.MigrateModel(db)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Println("Auto Migration Successfully")

	//gin.SetMode(gin.ReleaseMode) //enable when you want to production
	g := gin.Default()

	//Setting Config CORS
	c := cors.DefaultConfig()
	c.AllowWildcard = true
	c.AllowCredentials = true
	c.AllowOrigins = []string{"*"}
	//c.AllowOrigins = []string{"https://*.dbo.com"}
	c.AddAllowHeaders("Authorization", "Content-Type")
	c.AddExposeHeaders("Authorization", "Content-Type")

	//Use CORS
	g.Use(cors.New(c))

	//use Validation
	v := validator.New()

	//Inisialisasi Handler / Endpoint
	h := handlers.Context{Gin: g, DB: db, Validator: v}

	h.API("")

	Host := env.Get().AppHost
	Port := env.Get().AppPort

	Address := fmt.Sprintf("%s:%d", Host, Port)
	s := &http.Server{Addr: Address, Handler: g}

	log.Printf("start listen and serve at %s: %v", Host, Port)

	err = s.ListenAndServe()
	if err != nil {
		log.Fatal("failed to connect to server")
		return
	}
}
