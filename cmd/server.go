package main

import (
	"log"

	"github.com/izzatbey/soc-kafka-preprocess/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP ingest server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("🚀 Starting Go Preprocess Service on port %s...", cfg.Port)
		server.Start(cfg)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
