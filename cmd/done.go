/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/TuBl/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark a todo item as completed",
	Long:    ``,
	Run:     doneRun,
}

func updateItemStatus(item *todo.Item, toggle bool) {
	if toggle {
		item.ToggleDone()
		fmt.Printf("%q %v \n", item.Text, "toggle done status")
	} else {
		item.MarkDone()
		fmt.Printf("%q %v \n", item.Text, "marked done")
	}
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))

	if err != nil {
		log.Fatalln("Failed to read items from dataFile")
	}
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label \n", err)
	}

	toggle, err := cmd.Flags().GetBool("toggle")
	if err != nil {
		log.Fatalln("Error reading toggle flag:", err)
	}

	if i > 0 && i <= len(items) {
		updateItemStatus(&items[i-1], toggle)
		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)

	doneCmd.Flags().BoolP("toggle", "t", false, "Toggle the done status instead of marking done")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
