package main

import (
	"iris-psql/models"
	"log"

	"github.com/kataras/iris/v12"
)

// delete >> /customer/{id}
func (r *Repository) deleteCustomerData(ctx iris.Context) {
	customerModels := &models.Customers{}

	id := ctx.Params().Get("id")

	if id == "" {
		ctx.JSON(iris.Map{"message": "id cannot be empty"})
		log.Print("System: id cannot be empty")
		return
	}

	err := r.DB.Delete(customerModels, id)

	if err.Error != nil {
		ctx.JSON(iris.Map{"message": "could not delete customer data"})
		log.Print("System: could not delete customer data")
	} else {
		ctx.JSON(iris.Map{"message": "customer data deleted successfully"})
		log.Print("System: customer data deleted successfully")
	}
}
