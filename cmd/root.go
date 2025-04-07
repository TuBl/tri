/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo application",
	Long: `A simple yet elegant way to manager your todo list
	in the CLI`,
}

var dataFile string

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName("triconfig")

	home, err := os.UserHomeDir() // Correctly get the user's home directory
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	fmt.Println("home", home)

	configPath := home + "/.config" // Construct the full path
	viper.AddConfigPath(configPath) // Use the full path
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config:", viper.ConfigFileUsed(), err)
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {

	cobra.OnInitialize(initConfig)
	home, err := homedir.Dir()

	if err != nil {
		log.Println("Unable to detect home dire, pleasee set data file use --datafile")
	}

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+"tridos.json",
		"data file to store todos")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tri.json)")
}
