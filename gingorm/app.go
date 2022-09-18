package main

import (
	"slq.me/controller"
)

func main() {
	r := controller.GetBaseRouter()

	{
		controller.SetUserRouter(r)
	}

	r.Run(":8080")
}
