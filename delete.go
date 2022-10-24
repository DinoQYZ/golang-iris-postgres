package main

import (
	"fmt"
	"iris-psql/models"

	"github.com/kataras/iris/v12"
)

// delete >> /customer/{id}
func (r *Repository) deleteCustomerData(ctx iris.Context) {
	customerModels := &models.Customers{}

	id := ctx.Params().Get("id")

	if id == "" {
		ctx.JSON(iris.Map{"message": "id cannot be empty"})
		fmt.Println("System: id cannot be empty")
		return
	}

	err := r.DB.Delete(customerModels, id)

	if err.Error != nil {
		ctx.JSON(iris.Map{"message": "could not delete customer data"})
		fmt.Println("System: could not delete customer data")
	} else {
		ctx.JSON(iris.Map{"message": "customer data deleted successfully"})
		fmt.Println("System: customer data deleted successfully")
	}
}
