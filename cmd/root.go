package main

import (
	"fmt"
	"os"

	"github.com/izzatbey/soc-kafka-preprocess/internal/config"
	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "soc-preprocess",
	Short: "Log preprocessing service",
	Long:  "SOC Kafka Preprocess pipeline using Go.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cfg = config.Load()
}
