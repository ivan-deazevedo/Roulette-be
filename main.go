package main

import "lunch-be/app"

func main() {
	var a app.App
	a.CreateConnection()
	a.Routes()
	a.Run()
}
