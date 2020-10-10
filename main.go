package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"

	"github.com/a-berahman/auth-api/common"
	"github.com/a-berahman/auth-api/handler"
	"github.com/apex/gateway"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "auth-api", log.LstdFlags)
	authHandler := handler.NewAuthHandler(l)
	keyHandler := handler.NewKeyhHandler(l)

	keyHandler.GenerateKey()

	mx := mux.NewRouter()
	//set pprof
	mx.PathPrefix("/debug/").Handler(http.DefaultServeMux)

	postRouter := mx.Methods(http.MethodPost).Subrouter()

	postRouter.HandleFunc("/v1/auth/token", authHandler.GetToken)
	postRouter.HandleFunc("/v1/auth/validate", authHandler.ValidateToken)

	getRouter := mx.Methods(http.MethodGet).Subrouter()

	//SWAGGER CONFIGURATION
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	//*********************************************************************

	port := os.Getenv(common.PORT)

	serv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mx,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		l.Printf("application was ran in %s mode on port: %s \n ", os.Getenv(common.MODE), port)
		if os.Getenv(common.MODE) == common.DEV_MODE {
			serv.ListenAndServe()
		} else if os.Getenv(common.MODE) == common.PRD_MODE {
			gateway.ListenAndServe(fmt.Sprintf(":%s", port), nil)
		}
	}()
	sigChan := make(chan os.Signal)
	//registers the given channel to receive notifications of the specified signals
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recived terminate, graceful shutdown", sig)
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// serv.Shutdown(ctx)
}
