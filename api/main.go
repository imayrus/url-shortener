package main

import(
	"github.com/imayrus/url-shortener/database"
	"github.com/imayrus/url-shortener/setuproutes"
)

func main (){
	database.DatabaseSetup()
	routes.SetupAndListen()

}