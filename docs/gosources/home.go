/*
	@ Date of creation :
*/

package main

import "github.com/kataras/iris"

import "fmt"

func main() {

	home := func(ctx *iris.Context) {
		fmt.Println("Welcome to home...")
		ctx.SetContentType("application/json")
		ctx.JSON(iris.StatusOK, iris.Map{"name": "Rishikesh", "age": "24"})
	}

	iris.Post("/v1", home)
	iris.Listen(":8080")
}
