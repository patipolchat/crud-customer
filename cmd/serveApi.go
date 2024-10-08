/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crud-customer/config"
	"crud-customer/internal/http"
	"crud-customer/pkg/database"
	"crud-customer/pkg/server"
	"crud-customer/util"
	"github.com/spf13/cobra"
)

// serveApiCmd represents the serveApi command
var serveApiCmd = &cobra.Command{
	Use:   "serveApi",
	Short: "Start the API server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := util.GetConfig[config.Config]()
		if err != nil {
			panic(err)
		}
		db, err := database.NewGormDB(cfg)
		if err != nil {
			panic(err)
		}
		if err := db.AutoMigrate(); err != nil {
			panic("failed to automigrate")
		}
		serv := server.NewServer(cfg)
		app := http.NewApp(cfg, db, serv)
		app.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveApiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveApiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveApiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
