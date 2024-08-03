package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/themethaithian/go-pos-system/app"
	"github.com/themethaithian/go-pos-system/app/authen"
	"github.com/themethaithian/go-pos-system/app/product"
	"github.com/themethaithian/go-pos-system/config"
	"github.com/themethaithian/go-pos-system/database"
	"github.com/themethaithian/go-pos-system/middleware"
)

func main() {
	router := app.NewRouterHTTP()

	postgres := database.NewPostgres()
	validator := validator.New()

	authenStorage := authen.NewStorage(postgres)
	authenHandler := authen.NewHandler(authenStorage)

	router.POST("/login", authenHandler.Login)

	mdw := middleware.New()
	router.Use(mdw.VerifyToken)

	productStorage := product.NewStorage(postgres)
	productHandler := product.NewHandler(productStorage, validator)

	router.POST("/add-product", productHandler.NewProduct)
	router.POST("/edit-product/{id}", productHandler.EditProduct)

	server := http.Server{
		Addr:    ":" + config.Val.Port,
		Handler: router,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		d := time.Duration(5 * time.Second)
		fmt.Printf("shutting down init %s ...", d)
		// We received an interrupt signal, shut down.
		ctx, cancel := context.WithTimeout(context.Background(), d)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			fmt.Printf("HTTP server Shutdown: %s", err.Error())
		}
		close(idleConnsClosed)
	}()

	fmt.Println(":" + config.Val.Port + " is serve")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("HTTP server ListenAndServe: %s", err.Error())
		return
	}

	<-idleConnsClosed
	fmt.Println("gracefully")
}
