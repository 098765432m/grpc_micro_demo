package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/098765432m/grpc-micro-demo/hotel-service/consts"
	"github.com/098765432m/grpc-micro-demo/hotel-service/internal/server"
	"github.com/098765432m/grpc-micro-demo/hotel-service/scripts/pb"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have run hotel service com!mand")

		router := gin.Default()

		//Test gin
		router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// Set up gRpc
		go startGrpcServices()

		// Tien to /v1 url
		v1 := router.Group("/v1")

		// Function Set up Router
		setUpRouters(v1)

		// Run gin router
		if err := router.Run(fmt.Sprintf(":%d", consts.API_PORT)); err != nil {
			log.Fatal(err)
		}
	},
}

func setUpRouters(router *gin.RouterGroup) {

}

// khoi tao grpc server
func startGrpcServices() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", consts.GRPC_SERVER_PORT))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHotelServiceServer(grpcServer, &server.HotelServer{})

	// Connect to hotel service start
	// conn, err := grpc.NewClient("localhost:50022")
	// if err != nil {
	// 	log.Fatal("failed to connect to hotel service: ", err)
	// }

	// defer conn.Close()

	// Connect to hotel service end

	fmt.Printf("User gRpc running on port: %d\n", consts.GRPC_SERVER_PORT)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Cannot run user grpc server")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("failed to run hotel service command: ", err)
	}
}
