package cmd

import (
	"os"

	"github.com/silva4dev/banking-pix-enterprise-grade-clean-project/application/grpc"
	"github.com/silva4dev/banking-pix-enterprise-grade-clean-project/infrastructure/db"
	"github.com/spf13/cobra"
)

var portNumber int

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC Server port")
}
