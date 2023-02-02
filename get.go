package main

import (
	"fmt"
	"iris-psql/models"
	"log"

	"github.com/kataras/iris/v12"
)

// get >> /customer/
func (r *Repository) getCustomerData(ctx iris.Context) {
	customerModels := &[]models.Customers{}

	err := r.DB.Find(customerModels).Error
	if err != nil {
		ctx.JSON(iris.Map{"message": "fail to get customer data"})
		log.Print("System: fail to get customer data")
	} else {
		ctx.JSON(iris.Map{"data": customerModels})
		log.Print("System: customer data fetched successfully")
	}
}

// get >> /customer/{id}
func (r *Repository) getTargetCustomerData(ctx iris.Context) {
	customerModels := &models.Customers{}

	id := ctx.Params().Get("id")
	fmt.Printf("System: get target id [%v]\n", id)

	if id == "" {
		ctx.JSON(iris.Map{"message": "id cannot be empty"})
		log.Print("System: id cannot be empty")
		return
	}

	err := r.DB.Where("ID = ?", id).First(customerModels).Error

	if err != nil {
		ctx.JSON(iris.Map{"message": "could not get the customer data"})
		log.Print("System: could not get the customer data")
	} else {
		ctx.JSON(iris.Map{"data": customerModels})
		log.Print("System: customer id fetched successfully")
	}
}
