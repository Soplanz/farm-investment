package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"farm-investment/internal/config"
	"farm-investment/internal/database"
	"farm-investment/internal/handler"
	"farm-investment/internal/repository"
	"farm-investment/internal/routes"
	"farm-investment/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "rest",
	Short: "Run the API server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}
		log.Println("Config loaded")

		// connect DB
		db, err := database.ConnectDB(cfg)
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}
		defer db.Close()
		log.Println("Database connected")

		r := gin.Default()

		userRepo := repository.NewUserRepository(db)
		userService := service.NewUserService(userRepo)
		userHandler := handler.NewUserHandler(userService)

		routes.RegisterRoutes(r, userHandler)

		srv := &http.Server{
			Addr:    ":8080",
			Handler: r,
		}

		// run server
		go func() {
			log.Println("Server running on :8080")
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Fail to listen: %s\n", err)
			}
		}()

		// graceful shutdown
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}

		log.Println("Server exited properly")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
