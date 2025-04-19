/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TuBl/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for specific text in the todo list",
	Long:  ``,
	Run:   searchList,
}

func searchList(cmd *cobra.Command, args []string) {
	rawPath := viper.GetString("datafile")
	expandedPath := os.ExpandEnv(rawPath)

	if len(args) == 0 {
		fmt.Println("Please provide a search term")
		return
	}
	query := args[0]
	items, err := todo.ReadItems(expandedPath)

	if err != nil {
		log.Printf("%v", err)
	}

	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Text), strings.ToLower((query))) {
			fmt.Printf("- %s [Priority: %d, Done: %v]\n", item.Text, item.Priority, item.Done)
		}
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
