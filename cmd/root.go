package cmd

import (
	"fmt"
	"os"

	servicebus "github.com/Azure/azure-service-bus-go"
	"github.com/spf13/cobra"
)

var queueName string
var messageCount int
var queue *servicebus.Queue

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kedabus",
	Short: "used to pruduce/consume service bus messages",
	Long:  `used to pruduce/consume service bus messages.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		// make sure conn string exists for all child commands
		connStr := os.Getenv("SERVICEBUS_CONNECTION_STRING")
		if connStr == "" {
			return fmt.Errorf("FATAL: expected environment variable SERVICEBUS_CONNECTION_STRING not set")
		}

		// Create a client to communicate with a Service Bus Namespace.
		ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(connStr))
		if err != nil {
			return err
		}

		// Create a client to communicate with the queue. (The queue must have already been created)
		queue, err = ns.NewQueue(queueName)
		if err != nil {
			return err
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&queueName, "queue", "tasks", "queue to use")
	rootCmd.PersistentFlags().IntVarP(&messageCount, "count", "n", 1, "number of messages to produce or consume")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
