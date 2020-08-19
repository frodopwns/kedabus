package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	servicebus "github.com/Azure/azure-service-bus-go"
	"github.com/spf13/cobra"
)

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "produce messages",
	Long:  `produce {count} messages in {queue}. This is a utility command for testing the consumer.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("produce called")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// create {count} messages in the service bus
		for i := 1; i <= messageCount; i++ {
			log.Println("sending message", i)
			err := queue.Send(ctx, servicebus.NewMessageFromString(fmt.Sprintf("Hello, World #%d !!!", i)))
			if err != nil {
				log.Fatal(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(produceCmd)
}
