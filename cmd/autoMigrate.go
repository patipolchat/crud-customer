/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crud-customer/config"
	"crud-customer/pkg/database"
	"crud-customer/util"
	"fmt"
	"github.com/spf13/cobra"
)

// autoMigrateCmd represents the autoMigrate command
var autoMigrateCmd = &cobra.Command{
	Use:   "autoMigrate",
	Short: "Auto migrate will create the table in the database",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := util.GetConfig[config.Config]()
		if err != nil {
			panic("failed to get config")
		}

		db, err := database.NewGormDB(cfg)
		if err != nil {
			panic("failed to connect database")
		}

		if err := db.AutoMigrate(); err != nil {
			panic("failed to automigrate")
		}
		fmt.Println("AutoMigrate success")
	},
}

func init() {
	rootCmd.AddCommand(autoMigrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoMigrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// autoMigrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
