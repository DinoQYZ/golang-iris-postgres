package main

import (
	"fmt"
	"iris-psql/models"

	"github.com/kataras/iris/v12"
)

// get >> /customer/
func (r *Repository) getCustomerData(ctx iris.Context) {
	customerModels := &[]models.Customers{}

	err := r.DB.Find(customerModels).Error
	if err != nil {
		ctx.JSON(iris.Map{"message": "fail to get customers data"})
		fmt.Println("System: fail to get customers data")
	} else {
		ctx.JSON(iris.Map{"data": customerModels})
		fmt.Println("System: customers data fetched successfully")
	}
}

// get >> /customer/{id}
func (r *Repository) getTargetCustomerData(ctx iris.Context) {
	customerModels := &models.Customers{}

	id := ctx.Params().Get("id")
	fmt.Printf("System: get target id [%v]\n", id)

	if id == "" {
		ctx.JSON(iris.Map{"message": "id cannot be empty"})
		fmt.Println("System: id cannot be empty")
		return
	}

	err := r.DB.Where("ID = ?", id).First(customerModels).Error

	if err != nil {
		ctx.JSON(iris.Map{"message": "could not get the customer data"})
		fmt.Println("System: could not get the customer data")
	} else {
		ctx.JSON(iris.Map{"data": customerModels})
		fmt.Println("System: customer id fetched successfully")
	}
}
