package main

import (
	"log"

	"github.com/kataras/iris/v12"
)

func defaultGreet(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Welcome to customer system"})
	log.Print("System: greeting message sent")
}

func (r *Repository) firstData() {
	customer := Customer{
		Name:  "winter",
		Age:   21,
		Email: "winter@gm.com",
	}

	r.DB.Create(&customer)

	customer = Customer{
		Name:  "summer",
		Age:   22,
		Email: "summer@gm.com",
	}

	r.DB.Create(&customer)
}
