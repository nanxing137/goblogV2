package main

import (
	"fmt"
	v1 "goblogV2/routers/api/v1"

	"goblogV2/routers"
)

func main() {
	fmt.Println("open:", "http://localhost:8080/")

	r := routers.InitRouter()
	v1.Init(r)

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
