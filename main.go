package main

import (
	"fmt"
	"os"
	"template-api-gecho/constant"
	"template-api-gecho/routes"
)

// Starting server
// godotenv loaded at utils/log.go
// because log.go is loaded first than main init
// TODO: Change Architecture
func main() {
	echo := routes.Routing.GetRoutes(routes.Routing{})

	address := os.Getenv(constant.APPHost)
	port := os.Getenv(constant.APPPort)
	host := fmt.Sprintf("%s:%s", address, port)

	err := echo.Start(host)
	fmt.Println(err)
}
