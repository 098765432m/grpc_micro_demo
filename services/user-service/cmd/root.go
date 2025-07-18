package cmd

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have run user service")

		// Start Gin
		router := gin.Default()

		router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1 := router.Group("/v1")

		setUpRouters(v1)

		if err := router.Run(fmt.Sprintf(":%d", 3101)); err != nil {
			log.Fatal("failed to run REST server: ", err)
		}
	},
}

func setUpRouters(router *gin.RouterGroup) {
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("failed to start user service command: ", err)
	}
}
