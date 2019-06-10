package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	"ttany-chat-service/chat"
	"ttany-chat-service/middlewares"
	"ttany-chat-service/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.Infoln("Server initializing...")
	ConfigRuntime()
	utils.LoadConfig()
	utils.LoadLogConfig()

	db := utils.InitDB()
	defer db.Close()

	InitForDebug()

	r := gin.New()

	r.Use(middlewares.LoggerMiddleware())
	r.Use(gin.Recovery())

	api := r.Group("api")
	v1 := api.Group("v1")
	chat.ChatRoutes(v1)

	srv := &http.Server{
		Addr:    viper.GetString("server.port"),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func InitForDebug() {
	db := utils.GetDB()
	if db.DropTable(&chat.ParticipantModel{}, &chat.RoomModel{}).Error != nil {
		log.Error("Patlaidkkkk")
	}
	db.AutoMigrate(&chat.ParticipantModel{}, &chat.RoomModel{})

	room := chat.RoomModel{
		RoomID:  "2671c20b-0b09-4648-8f4c-0369b284e9b4",
		AdminID: "8fcc9a26-04d0-4f40-8eaf-3d705669acf6",
		Name:    "Sohbet muhabbet",
		Type:    "direct",
		Status:  "active",
		ParticipantModels: []chat.ParticipantModel{
			{
				RoomID: "2671c20b-0b09-4648-8f4c-0369b284e9b4",
				UserID: "8fcc9a26-04d0-4f40-8eaf-3d705669acf6",
			},
			{
				RoomID: "2671c20b-0b09-4648-8f4c-0369b284e9b4",
				UserID: "cb235d43-e300-4ba6-99be-390ce0812a85",
			},
		},
	}

	if result := db.Save(&room); result.Error != nil {
		log.Error("Error occured", result.Error)
	}
}
