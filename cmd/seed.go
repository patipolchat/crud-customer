/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crud-customer/config"
	"crud-customer/database"
	"crud-customer/util"
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed will seed the database with some data",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := util.GetConfig[config.Config]()
		if err != nil {
			panic("failed to get config")
		}

		db, err := database.NewGormDB(cfg)

		if err != nil {
			panic("failed to connect database")
		}

		if err := db.Seed(); err != nil {
			panic("failed to seed")
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
