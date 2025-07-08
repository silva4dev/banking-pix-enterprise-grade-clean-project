package cmd

import (
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/silva4dev/banking-pix-enterprise-grade-clean-project/application/kafka"
	"github.com/silva4dev/banking-pix-enterprise-grade-clean-project/infrastructure/db"
	"github.com/spf13/cobra"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		deliveryChan := make(chan ckafka.Event)
		database := db.ConnectDB(os.Getenv("env"))
		producer := kafka.NewKafkaProducer()
		go kafka.DeliveryReport(deliveryChan)
		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
