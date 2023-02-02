package main

import (
	"iris-psql/models"
	"iris-psql/storages"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type Customer struct {
	Name  string
	Age   uint64
	Email string
}

type Repository struct {
	DB *gorm.DB
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storages.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storages.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateCustomerData(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}

	app := iris.New()

	//default data
	r.firstData()

	//set up all routes
	r.SetupRoutes(app)

	app.Run(iris.Addr(":4000"))
}

func (r *Repository) SetupRoutes(app *iris.Application) {
	//default html
	app.RegisterView(iris.HTML("./views", ".html"))

	//default
	app.Get("/", defaultGreet)

	//get customer
	app.Get("/customer", r.getCustomerData)

	//get single user
	app.Get("/customer/{id}", r.getTargetCustomerData)

	//add customer
	app.Post("/customer", r.addCustomerData)

	//update customer
	app.Put("/customer/{id}", r.updateCustomerData)

	//delete customer
	app.Delete("/customer/{id}", r.deleteCustomerData)
}
