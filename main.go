package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const jsonPath = `C:\Users\tashr\Desktop\projects\golang\practice\json`
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
	ID   int `json:"taskId"`
	Message string `json:"taskInfo"`
	Completed bool `json:"Completed"`
}

var commandsMap = map[string]func(){
	"/taskCreate": createTask,
	"/taskDelete": deleteTask,
	"/listTasks": listTasks,
}

func main() {
	taskOptionsMessage()
}

func taskOptionsMessage(){
	fmt.Println("Hello golang!")
	fmt.Println("1. Create Task\n2. Delete Task\n3. List tasks currently in queue")
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
	case "3":
		commandsMap["/listTasks"]()
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
	removeErr := os.Remove("_")
	if removeErr != nil {
		fmt.Println("Error removing file:", removeErr)
	} else {
		fmt.Println("Temporary file deleted.")
	}
}

func getNumberOfFiles()(int){
	filecount := 0
	err := filepath.WalkDir(jsonPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
		fmt.Println("Error checking directory")
		return err
		}
		if !d.IsDir(){
			filecount++
		}
		return nil
	})
	if err != nil {
		fmt.Println(err) 
	}
	fmt.Printf("Number of files: %d", filecount)
	return filecount
}

func readFile(file string){
	data, readErr := os.ReadFile(file)
	if readErr != nil {
		fmt.Println("Error reading file:", readErr)
		return
	}
	fmt.Println("File contents:\n", string(data))
	if len(string(data)) == 0 {
		fmt.Println("File is empty")
	}
}

func deleteTask() {
	//var arrayPosition int
	fmt.Println("[+] Deleting task...")
}

func getTaskId(task string){

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
	fmt.Println("[+] Creating task struct... ")
	var (
		taskId int
		taskInfo string
	)

	taskName := getTaskName()
	if taskName == "" {
		fmt.Println("Task name cannot be empty.")
		return
	}

	taskId = getNumberOfFiles()
	taskInfo = addTaskInfo()
	fmt.Println("[+] Task being created...")

	defer formJsonStructure(taskName, taskId, taskInfo)
}

func formJsonStructure(taskName string, taskId int, taskInfo string){
	task := Task{
		Name: taskName,
		ID:   taskId,
		Message: taskInfo,
		Completed: false,
	}
	saveJson(task, taskName)
}

func saveJson(task Task, taskName string){
	fmt.Println("[+] Creating json data")
	jsonBytes, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling task:", err)
		return
	}
	fmt.Println("JSON data:\n", string(jsonBytes))
	saveJsonData(jsonBytes, taskName)
}

func listTasks()  {
	fmt.Println("[+] Attemting to list tasks")
	files,err:=os.ReadDir(jsonPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("List of tasks?? %s\n", files)
	for i:=0; i< len(files);i++{
		fmt.Printf("File %d: %s\n",i+1, files[i])
	}
}

func completeTheTask(){

}

func saveJsonData(jsonData []byte, taskName string){
	fmt.Println("[+] Saving to json... ")
	jsonFilePath := fmt.Sprintf("%s\\%s.json", jsonPath, strings.ReplaceAll(taskName, " ", "_"))
	fmt.Println(jsonFilePath)
	file, readerr := os.OpenFile(jsonFilePath,os.O_CREATE, 0644)
	if readerr != nil{
		fmt.Println("Error reading user data")
	}
	saveTofile(jsonData, file)
}

func saveTofile(jsonData []byte, file *os.File){
	_, wErr :=file.Write(jsonData)
	if wErr != nil {
		fmt.Println(wErr)
	}
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