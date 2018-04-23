package main

import "github.com/xzdbd/portal/internal/api"

func main() {
	r := api.GinAPIRouter()
	r.Run(":8080")
}
