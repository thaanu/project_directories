package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var parentProjectDirectory string

func main() {

	godotenv.Load(".env")

	fmt.Println("Wow! New Project")

	// Get the required details
	clientName := askForInput("Client Name: ")
	clientEmail := askForInput("Client Email: ")
	clientContact := askForInput("Client Contact: ")
	projectName := askForInput("Enter the project name: ")
	projectDate := askForInput("Enter Project Date (YYYY-MM-DD): ")

	parentProjectDirectory = getProjectDir() + "/" + projectName

	// First create the current project directory
	os.Mkdir(getProjectDir(), 0755)
	os.Mkdir(parentProjectDirectory, 0755)

	createProjectDirectories()

	createIndexFile(clientName, clientEmail, clientContact, projectName, projectDate)

	fmt.Println("Project Name: " + projectName)
	fmt.Println("Project Directory: " + parentProjectDirectory)

}

func getProjectDir() string {

	environmentValue, err := getEnv("PROJECT_DIR")
	if err != nil {
		homeDir, _ := os.UserHomeDir()
		return homeDir
	}
	return environmentValue
}

func getEnv(key string) (value string, err error) {

	f, _ := os.Getwd()
	projectDirectory := filepath.Dir(f) + "/" + filepath.Base(f)
	environmentFile := projectDirectory + "/.env"

	// Check if the environment file exist
	_, err = os.Stat(environmentFile)
	if err != nil {
		return
	}
	err = godotenv.Load(environmentFile)
	if err != nil {
		return
	}
	value = os.Getenv(key)
	return
}

func askForInput(question string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, question+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func createProjectDirectories() (response bool, err error) {

	subDirectoryNames := [5]string{"RAW Files", "Music", "Assets", "Project Files", "Deliverables"}

	for _, folder := range subDirectoryNames {
		actualDirectory := parentProjectDirectory + "/" + folder
		os.Mkdir(actualDirectory, 0755)
	}

	response = true

	return
}

func createIndexFile(clientName string, clientEmail string, clientContact string, projectName string, projectDate string) {
	inputText := "Client Name: " + clientName + "\n"
	inputText = inputText + "Client Email: " + clientEmail + "\n"
	inputText = inputText + "Client Contact: " + clientContact + "\n"
	inputText = inputText + "Project Name: " + projectName + "\n"
	inputText = inputText + "Project Date: " + projectDate

	os.WriteFile(parentProjectDirectory+"/readme.txt", []byte(inputText), 0644)

}
