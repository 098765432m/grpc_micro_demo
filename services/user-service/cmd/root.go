package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/098765432m/grpc-micro/user-service/consts"
	"github.com/098765432m/grpc-micro/user-service/internal/user/server"
	"github.com/098765432m/grpc-micro/user-service/scripts/pb"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {

		// Start Gin
		router := gin.Default()

		// test router
		router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1 := router.Group("/v1")

		setUpRouters(v1)

		// Run grpc server in background
		go startGrpcServer()

		if err := router.Run(fmt.Sprintf(":%d", consts.API_PORT)); err != nil {
			log.Fatal("failed to run REST server: ", err)
		}
	},
}

func setUpRouters(router *gin.RouterGroup) {
}

// Set up grpc server
func startGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", consts.GPRC_SERVER_PORT))
	if err != nil {
		log.Fatal("failed to start user service grpc server: ", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &server.UserServer{})

	fmt.Printf("User service grpc running on port: %d", consts.GPRC_SERVER_PORT)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to run user service grpc server: ", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("failed to start user service command: ", err)
	}
}
