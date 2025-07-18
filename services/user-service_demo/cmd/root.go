package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/098765432m/user-service/internal/handler"
	"github.com/098765432m/user-service/pb"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have run an user service command")

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
		if err := router.Run(fmt.Sprintf(":%d", 3101)); err != nil {
			log.Fatal(err)
		}

	},
}

// Diem de set router
func setUpRouters(router *gin.RouterGroup) {

}

// khoi tao grpc server
func startGrpcServices() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, &handler.UserHanlder{})

	// Connect to hotel service start
	// conn, err := grpc.NewClient("localhost:50022")
	// if err != nil {
	// 	log.Fatal("failed to connect to hotel service: ", err)
	// }

	// defer conn.Close()

	// Connect to hotel service end

	fmt.Printf("User gRpc running on port: %d\n", 50051)

	if err := server.Serve(lis); err != nil {
		log.Fatal("Cannot run user grpc server")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Khong the thuc chay lenh")
	}
}
