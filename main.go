package main

import (
	"fmt"
	"os"
	"os/exec"
)

func convertMp3ToAmr(inputFile, outputFile string) error {
	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file %s does not exist", inputFile)
	}

	// Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current working directory: %v", err)
	}

	// Prepare the Docker command to run ffmpeg inside the container
	cmd := exec.Command(
		"docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/data", dir), // Mount current directory to /data in the container
		"jrottenberg/ffmpeg:4.4-alpine", // Use the jrottenberg/ffmpeg image
		"-i", "/data/"+inputFile,        // Path inside the container
		"-ar", "8000", // Set audio sample rate
		"-ac", "1", // Set audio channels to 1 (mono)
		"-f", "amr", // Set output format to AMR
		"/data/"+outputFile, // Output file path inside the container
	)

	// Run the command and capture output
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	fmt.Printf("Converting %s to %s...\n", inputFile, outputFile)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg conversion failed: %v", err)
	}

	fmt.Println("Conversion completed successfully!")
	return nil
}

func main() {
	// Define the input and output file names
	inputFile := "input.mp3"   // Hardcoded input file name
	outputFile := "output.amr" // Hardcoded output file name

	// Perform conversion
	if err := convertMp3ToAmr(inputFile, outputFile); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
