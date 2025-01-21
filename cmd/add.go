/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Task  struct {
	ID  int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		tasks := loadTasks()
		tasks = append(tasks, Task{ID: len(tasks) + 1, Name: name, Date: time.Now().Format("2006-01-02") })
		addTask(tasks)
		fmt.Println("Task added:", name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func loadTasks() []Task {
	file, err := os.Open("output.json")
	if os.IsNotExist(err) {
	 return []Task{}
	} else if err != nil {
	 panic(err)
	}
	defer file.Close()
   
	var tasks []Task
	json.NewDecoder(file).Decode(&tasks)
	return tasks
   }

func addTask(tasks []Task) {
	var filename string = "output.json"

file, err := os.Create(filename)
if err != nil {
	log.Fatal(err)
}
defer file.Close()

json.NewEncoder(file).Encode(tasks)

}
