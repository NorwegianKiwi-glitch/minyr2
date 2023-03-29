package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Convert Celsius to Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func main() {
	// Define command-line flags
	convertFlag := flag.Bool("convert", false, "convert temperature data from Celsius to Fahrenheit")
	flag.Parse()

	if *convertFlag {
		// Check if output file already exists
		if _, err := os.Stat("kjevik-tempfahr-20220318-20230318.csv"); !os.IsNotExist(err) {
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
		inputFile, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
		if err != nil {
			fmt.Println("Error opening input file:", err)
			return
		}
		defer inputFile.Close()

		// Create a new scanner to read the input CSV file
		inputScanner := bufio.NewScanner(inputFile)

		// Create a new CSV writer to write the output CSV file
		outputFile, err := os.Create("kjevik-tempfahr-20220318-20230318.csv")
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
			fields[3] = strconv.FormatFloat(fahrenheit, 'f', 2, 64)

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