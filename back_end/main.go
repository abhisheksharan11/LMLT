package main

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/LMLT/back_end/middleware"
	"os"
	"github.com/urfave/negroni"

)

func main() {


	// Prints the address of our api to the console
	logrus.Info("Starting Server on http://localhost:", os.Getenv("PORT"))

	server := negroni.New()
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.LogMiddleware())

	// Attach app router
	appServer:=server.UseHandler(router.NewRouter())
	logrus.Fatal(http.ListenAndServe(":8888", appServer))
}
