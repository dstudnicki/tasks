/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [task id]",
	Short: "This command lets you complete a task.",
	Long: `This command lets you complete a task. 
	For example: ./task complete 1
	Then it completes the task by ID and saves it to output.json file with the completed task.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.Println("Error: Task ID must be a number")
			return
		}
		completeTask(id)
		fmt.Println("Task completed:", id) // 
	},
}


func init() {
	rootCmd.AddCommand(completeCmd)
}

func completeTask(taskID int){
	tasks := loadTasks()

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Done = true
			break
		}
	}

	addTask(tasks)
}
