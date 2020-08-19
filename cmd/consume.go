package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	servicebus "github.com/Azure/azure-service-bus-go"
	"github.com/spf13/cobra"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "consume from the queue",
	Long:  `consume from the queue provided`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consume called")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		for {

			err := queue.ReceiveOne(
				ctx,
				servicebus.HandlerFunc(func(ctx context.Context, message *servicebus.Message) error {
					log.Println(string(message.Data))
					return message.Complete(ctx)
				}),
			)
			if err != nil {
				if strings.Contains(err.Error(), "deadline exceeded") {
					log.Println("no message returned in time, could be none left")
					time.Sleep(5 * time.Second)
					continue
				}
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(consumeCmd)
}
