package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"portfolio-tracker/controller"
	"portfolio-tracker/repository"
	"portfolio-tracker/service"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	log.Println("Starting server")

	db, err := initDB()

	if err != nil {
		log.Fatalf("Unable to initialize db %v\n", err)
	}

	router := gin.Default()

	router.Use(cors.Default())



	transactionRepository := repository.NewTransactionRepository(db.DB)
	transactionService := service.NewTransactionService(transactionRepository)
	controller.NewController(&controller.Config{
		R: router,
		TransactionService: transactionService,
	})

	srv:= &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	go func(){
		//service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

//Graceful server shutdown

// Wait for interrupt signal to gracefully shutdown the server
// with a timeout of 5 seconds.
quit := make(chan os.Signal, 1)
// kill (no param) default send syscall.SIGTERM
// kill -2 is syscall.SIGNINT
// kill -9 is syscall.SIGKILL but can't be catch, don't need to add it then
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit
log.Println("Shutdown Server...")

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)


if err := db.close(); err != nil {
	log.Fatalf("A problem occured gracefully shutting down the db connection: %v\n", err)
}

defer cancel()
if err := srv.Shutdown(ctx); err != nil {
	log.Fatal("Server shutdown", err)
}
// catching ctx.Done(). timeout of 5s.
<-ctx.Done()
log.Println("timeout of 5s")
log.Println("Server exiting")
}
