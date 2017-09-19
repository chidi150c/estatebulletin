package main

import (
"os"
)

func main(){
	port := os.Getenv("PORT")
	appHandler := app.NewAppHandler()
	handler := http_api.NewMyHandler(appHandler)
	server:= http_api.NewServer(port, handler)
	err := server.Open()
}