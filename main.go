package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			// Scanner encountered an error or EOF
			break
		}
		command := scanner.Text()

		switch strings.ToLower(command) {
		case "minyr":
			// Kill all foreground processes
			cmd := exec.Command("pkill", "-f", "")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to kill foreground processes: %v\n", err)
			}
		case "q", "quit":
			// Exit the loop and terminate the program
			fmt.Println("Exiting...")
			return
		default:
			// Unknown command, print an error message
			fmt.Println("Unknown command:", command)
		}
	}
	// Define command-line flags
	convertFlag := flag.Bool("convert", false, "convert temperature data from Celsius to Fahrenheit")
	flag.Parse()

	if *convertFlag {
		// Check if output file already exists
		if _, err := os.Stat("kjevik-temp-celsius-20220318-20230318.csv"); !os.IsNotExist(err) {
			// Output file already exists, prompt user to regenerate
			var regenerate string
			fmt.Print("Output file already exists. Regenerate? (j/n): ")
			fmt.Scanln(&regenerate)
			if regenerate != "j" && regenerate != "J" {
				fmt.Println("Exiting without generating new file.")
				return
			}
		}

		// Open the input CSV file
		inputFile, err := os.Open("yr/kjevik-temp-celsius-20220318-20230318.csv")
		if err != nil {
			fmt.Println("Error opening input file:", err)
			return
		}
		defer inputFile.Close()

		// Create a new scanner to read the input CSV file
		inputScanner := bufio.NewScanner(inputFile)

		// Create a new CSV writer to write the output CSV file
		outputFile, err := os.Create("yr/kjevik-temp-fahr-20220318-20230318.csv")
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputFile.Close()

		outputWriter := csv.NewWriter(outputFile)
		defer outputWriter.Flush()

		// Skriv ut den første linjen fra inputfilen til outputfilen
		if inputScanner.Scan() {
			firstLine := inputScanner.Text()
			if err = outputWriter.Write(strings.Split(firstLine, ";")); err != nil {
				fmt.Println("Error writing first line:", err)
				return
			}
		}

		// Loop through each line of the input CSV file
		for inputScanner.Scan() {
			// Split the line into fields
			fields := strings.Split(inputScanner.Text(), ";")

			// Check that the fields slice has at least 4 elements
			if len(fields) < 4 {
				fmt.Println("Error: Invalid input format.")
				continue
			}

			// Extract the last digit from the fourth column
			temperature, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				fmt.Println("Error parsing temperature:", err)
				continue
			}
			lastDigit := temperature - float64(int(temperature/10))*10

			// Convert Celsius to Fahrenheit
			fahrenheit := celsiusToFahrenheit(lastDigit)

			// Replace the temperature in the fourth column with the converted value
			temperatureString := strconv.FormatFloat(fahrenheit, 'f', 2, 64)
			temperatureParts := strings.Split(temperatureString, ".")
			fields[3] = temperatureParts[0] + "." + string(temperatureParts[1][0])

			// Write the updated line to the output CSV file
			err = outputWriter.Write(fields)
			if err != nil {
				fmt.Println("Error writing line to output file:", err)
				continue
			}
		}

		dataText := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Simon Helgen"
		err = outputWriter.Write([]string{dataText})
		if err != nil {
			fmt.Println("Error writing data text to output file:", err)
			return
		}
		fmt.Println("Temperature conversion complete.")
		return
	}
	// Wait for user input
	fmt.Println("Press enter to exit.")
	fmt.Scanln()
}
