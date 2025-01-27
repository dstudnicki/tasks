/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "This command lists all the tasks",
	Long: `This command lists all the tasks.
	For example: ./task list
	Then it lists all the tasks in the output.json file in tabular format.`,
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
		
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTasks() []Task { 
	writer := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	tasks := loadTasks()
	writer.Write(
		[]byte("ID\tTask\tCreated\tDone\n"),
	)
	for _, task := range tasks {
		writer.Write(
			[]byte(fmt.Sprintf("%d\t%s\t%s\t%t\n", task.ID, task.Name, task.Date, task.Done)),
		)
	}
	writer.Flush()
	return tasks	

}
