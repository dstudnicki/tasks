/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Task  struct {
	ID  int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
	Done bool   `json:"done"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "This command lets you add a task to the list",
	Long: `This command lets you add a task to the list. 
	For example: ./task add "Learn Go"
	Then it creates output.json file with the task added.
	`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.Join(args, " ")
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

