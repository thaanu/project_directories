package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)



func main() {

	godotenv.Load(".env")

	fmt.Println("Wow! New Project")

	fmt.Println(getEnv("PROJECT_DIR"))

	// Get the required details
	projectName := askForInput("Enter the project name: ")

	fmt.Println("Project Name: " + projectName)

}

func getEnv(value string) string {
	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(value)
}

func askForInput(question string) string {
	var s string
	r:=bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, question+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}