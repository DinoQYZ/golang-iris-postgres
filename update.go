package main

import (
	"iris-psql/models"
	"log"
	"strconv"

	"github.com/kataras/iris/v12"
)

// put >> /customer/{id}
func (r *Repository) updateCustomerData(ctx iris.Context) {

	id := ctx.Params().Get("id")
	inputName := ctx.PostValue("name")
	inputAge, _ := strconv.ParseUint(ctx.PostValue("age"), 10, 64)
	inputEmail := ctx.PostValue("email")

	customerModels := &models.Customers{
		Name:  inputName,
		Age:   inputAge,
		Email: inputEmail,
	}

	log.Printf("System: get target id [%v]\n", id)

	if id == "" {
		ctx.JSON(iris.Map{"message": "id cannot be empty"})
		log.Print("System: id cannot be empty")
		return
	}

	err := r.DB.Where("ID = ?", id).UpdateColumns(customerModels).First(customerModels).Error

	if err != nil {
		ctx.JSON(iris.Map{"message": "could not get the customer data"})
		log.Print("System: could not get the customer data")
	} else {
		ctx.JSON(iris.Map{"data": customerModels})
		log.Print("System: customer id fetched successfully")
	}
}
