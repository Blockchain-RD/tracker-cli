/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"time"

	"tracker-cli/requests"
	"tracker-cli/types"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		addTransaction(cmd)
	},
}

func addTransaction(cmd *cobra.Command) {
	id, err := cmd.Flags().GetString("id")
	if err != nil || id == "" {
		log.Fatalln("missing id value")
	}

	coin, err := cmd.Flags().GetString("coin")
	if err != nil || coin == "" {
		log.Fatalln("missing coin value")
	}

	date, err := cmd.Flags().GetString("date")
	if err != nil {
		log.Fatalln("missing date value")
	}

	parseDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatalln("invalid date format")
	}

	value, err := cmd.Flags().GetFloat64("value")
	if err != nil || value == 0 {
		log.Fatalln("missing quantity value")
	}

	newTransaction := types.Transaction{
		Id:    id,
		Coin:  coin,
		Date:  parseDate,
		Value: value,
	}

	err = requests.AddTransaction(&newTransaction)
	if err != nil {
		log.Fatalln("failed saving transaction")
	}
}

func init() {
	addCmd.Flags().String("id", "", "")
	addCmd.Flags().String("date", "", "")
	addCmd.Flags().String("coin", "", "")
	addCmd.Flags().Float64("value", 0, "")

	transactionCmd.AddCommand(addCmd)

}
