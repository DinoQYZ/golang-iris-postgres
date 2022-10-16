package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func defaultGreet(ctx iris.Context) {
	fmt.Println("System: greeting message sent")
	ctx.JSON(iris.Map{"message": "Welcome customer system"})
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
