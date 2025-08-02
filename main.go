package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const jsonDataPath  = `C:\Users\tashr\Desktop\projects\golang\pratice\json\userData.json`
var reader = bufio.NewReader(os.Stdin)

type OptionError struct {
	msg string
}

type TaskList struct {
	Name  string `json:"List name"`
	ID    string `json:"list id"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name string `json:"taskName"`
	ID   string `json:"taskId"`
	Message string `json:"taskInfo"`
}

var commandsMap = map[string]func(){
	"/taskCreate": createTask,
	"/taskDelete": deleteTask,
}

func main() {
	fmt.Println("Hello golang!")
	openUserInfo()
	fmt.Println("1. Create Task\n2. Delete Task")
	fmt.Print("Enter choice: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		if optionError, ok := err.(*OptionError); ok && optionError.IsInvalidOption(){
			fmt.Println("Invalid option chosen")
			main()
		}
	}

	choice := strings.TrimSpace(input)

	switch choice {
	case "1":
		commandsMap["/taskCreate"]()
	case "2":
		commandsMap["/taskDelete"]()
	default:
		fmt.Println("Invalid option")
	}
}

func (e *OptionError) Error() string {
	return e.msg
}

func (e *OptionError) IsInvalidOption() bool {
	return strings.Contains(e.msg, "invalid option")
}

func deleteUserInfo(){
	// Delete the file
	removeErr := os.Remove(jsonDataPath)
	if removeErr != nil {
		fmt.Println("Error removing file:", removeErr)
	} else {
		fmt.Println("Temporary file deleted.")
	}
}

func openUserInfo(){
	done := make(chan bool)
	// Checking for data from file
	_, err := os.Stat(jsonDataPath)
	if err != nil {
		go func() {
			createUserJson()
			done <- true
		}()
		<-done
	}
	readFile()
}

func readFile(){
	data, readErr := os.ReadFile(jsonDataPath)
	if readErr != nil {
		fmt.Println("Error reading file:", readErr)
		return
	}
	fmt.Println("File contents:\n", string(data))
	if len(string(data)) == 0 {
		fmt.Println("File is empty")
	}
}

func createUserJson(){
	file, createErr := os.Create(jsonDataPath)
	if createErr != nil {
		fmt.Println("Error creating file:", createErr)
		return
	}
	file.Close()
}

func deleteTask() {
	fmt.Println("[+] Deleting task...")
}

func getTaskName() string {
	fmt.Print("Name of task? ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading task name:", err)
		return ""
	}
	taskName := strings.TrimSpace(input)
	fmt.Println("This is the input:", taskName)
	return taskName
}

func createTask() {
	var (
		taskId string
		taskInfo string
	)

	taskName := getTaskName()
	if taskName == "" {
		fmt.Println("Task name cannot be empty.")
		return
	}

	taskId = fmt.Sprintf("%s.1", taskName)
	var taskIdRemovedSpaces  = removeLineSpacing(taskId)
	taskInfo = addTaskInfo()
	fmt.Println("[+] Task being created...")

	defer formJsonStructure(taskName, taskIdRemovedSpaces, taskInfo)
}

func formJsonStructure(taskName string, taskId string, taskInfo string){
	task := Task{
		Name: taskName,
		ID:   taskId,
		Message: taskInfo,
	}

	saveAsJson(task)
}

func saveAsJson(task Task){
	jsonBytes, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling task:", err)
		return
	}
	fmt.Println("JSON data:\n", string(jsonBytes))
}

func addTaskInfo() string {
	var taskInfo string
	fmt.Print("What would you like to add? \n")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading task name:", err)
		return ""
	}
	
	taskInfo = strings.TrimSpace(input)
	return taskInfo
}

func numberOfTasks() {

}

func removeLineSpacing(input string) string {
    noNewLines := strings.ReplaceAll(input, "\r\n", "")
    noNewLines = strings.ReplaceAll(noNewLines, "\n", "")
    return noNewLines
}