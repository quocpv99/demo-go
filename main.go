package main

import "go_api/routes"

func main() {
	r := routes.SetupBlogRoutes()
	r.Run(":8080")

}
