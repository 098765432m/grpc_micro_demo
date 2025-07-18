package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/098765432m/hotel-service/internal/handler"
	"github.com/098765432m/hotel-service/pb"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have run an hotel service command")

		router := gin.Default()

		//Test gin
		router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// Set up gRpc
		go StartGrpcServices()

		// Tien to /v1 url
		v1 := router.Group("/v1")

		// Function Set up Router
		setUpRouters(v1)

		fmt.Println("Hotel server running on port: 3102")

		// Run gin router
		if err := router.Run(fmt.Sprintf(":%d", 3102)); err != nil {
			log.Fatal("err")
		}

	},
}

// Diem de set router
func setUpRouters(router *gin.RouterGroup) {

}

// khoi tao grpc server
func StartGrpcServices() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterHotelServiceServer(server, &handler.HotelHanlderImpl{})

	fmt.Printf("User gRpc running on port: %d\n", 50052)

	if err := server.Serve(lis); err != nil {
		log.Fatal("Cannot run user grpc server")
	}
}

// func ConnectKafka() {
// 	topic := "hotel-topic"
// 	partition := 0

// 	// Connect to kafka
// 	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
// 	if err != nil {
// 		log.Fatal("failed to dial leader: ", err)
// 	}

// 	// Set timeout for Write operation
// 	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

// 	// Write message
// 	_, err = conn.WriteMessages(
// 		kafka.Message{Value: []byte("one!")},
// 		kafka.Message{Value: []byte("Two!")},
// 		kafka.Message{Value: []byte("three!")},
// 	)

// 	if err != nil {
// 		log.Fatal("failed to write messages: ", err)
// 	}

// 	if err := conn.Close(); err != nil {
// 		log.Fatal("failed to close writer: ", err)
// 	}
// }

func ReadKafkaMessage() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092", "localhost:9093"},
		Topic:    "user-topic",
		GroupID:  "group-a",
		MaxBytes: 10e6,
	})

	// Read start at id 42
	r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func WriteKafkaMessage() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "user-topic",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Khong the thuc chay lenh")
	}
}
