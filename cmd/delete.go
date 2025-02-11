/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "This command lets you delete a task.",
	Long: `This command lets you delete a task. 
	For example: ./task delete 1
	Then it delete the task by ID and saves the output.json file without the completed task.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.Println("Error: Task ID must be a number")
			return
		}
		deleteTask(id)
		fmt.Println("Task deleted:", id) //
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteTask(taskID int) {
	tasks := loadTasks()

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	addTask(tasks)
}
