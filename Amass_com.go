package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the domain: ")
	domain, _ := reader.ReadString('\n')
	domain = strings.TrimSpace(domain)

	fmt.Print("Enter the path to the config file: ")
	configFile, _ := reader.ReadString('\n')
	configFile = strings.TrimSpace(configFile)

	// Amass command with desired flags
	cmd := exec.Command("amass", "enum", "-d", domain, "-config", configFile, "-ip", "-brute", "-active", "-min-for-recursive", "2", "-o", "output.txt")

	// Run the Amass command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running Amass command: %v\n", err)
		fmt.Printf("Output: %s\n", output)
		log.Fatal("Exiting due to error.")
	}

	fmt.Printf("Amass output: %s\n", output)
}
