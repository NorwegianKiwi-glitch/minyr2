package yr

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	/*"github.com/NorwegianKiwi-glitch/funtemps/conv"*/)

// konverterer temperatur fra Celsius til Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func Convert() error {
	// Check if output file already exists
	if _, err := os.Stat("yr/kjevik-temp-fahr-20220318-20230318.csv"); !os.IsNotExist(err) {
		// Output file already exists, prompt user to regenerate
		var regenerate string
		fmt.Print("Output file already exists. Regenerate? (y/n): ")
		fmt.Scanln(&regenerate)
		if regenerate != "y" && regenerate != "Y" {
			fmt.Println("Exiting without generating new file.")
			return nil
		}
	}

	// Open the input CSV file
	inputFile, err := os.Open("yr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Error opening input file:", err)
	}
	defer inputFile.Close()

	// Create a new scanner to read the input CSV file
	inputScanner := bufio.NewScanner(inputFile)

	// Create a new CSV writer to write the output CSV file
	outputFile, err := os.Create("yr/kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Error creating output file:", err)
	}
	defer outputFile.Close()

	outputWriter := csv.NewWriter(outputFile)
	defer outputWriter.Flush()

	// Prints out the first line of the input CSV file
	if inputScanner.Scan() {
		firstLine := inputScanner.Text()
		if err = outputWriter.Write(strings.Split(firstLine, ";")); err != nil {
			fmt.Println("Error writing first line:", err)
		}
	}

	// Loop through each line of the input CSV file
	for inputScanner.Scan() {
		// Split the line into fields
		fields := strings.Split(inputScanner.Text(), ";")

		// Check that the fields slice has at least 4 elements
		if len(fields) != 4 {
			fmt.Println("Error: Invalid input format.")
			continue
		}

		// Extract the last digit from the fourth column
		temperatureField := fields[3]
		if temperatureField == "" {
			fmt.Println("Error: Temperature value is empty.")
			continue
		}
		temperature, err := strconv.ParseFloat(temperatureField, 64)
		if err != nil {
			fmt.Println("Error parsing temperature:", err)
			continue
		}

		if math.IsNaN(temperature) {
			fmt.Println("Error: Temperature is not a valid float64 value.")
			continue
		}

		lastDigit := temperature - float64(int(temperature/10))*10

		// Convert Celsius to Fahrenheit
		// fahrenheit := conv.CelsiusToFarenheit(lastDigit) //
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

	dataText := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Simon Helgen,,,"
	err = outputWriter.Write([]string{dataText})
	if err != nil {
		fmt.Println("Error writing data text to output file:", err)

	}

	return nil
}

func Average(unit string) (float64, error) {
	var filename string
	var tempColumn int
	var delimiter rune

	if unit == "c" {
		filename = "yr/kjevik-temp-celsius-20220318-20230318.csv"
		tempColumn = 3
		delimiter = ';'
	} else if unit == "f" {
		filename = "yr/kjevik-temp-fahr-20220318-20230318.csv"
		tempColumn = 3
		delimiter = ','
	} else {
		return 0, fmt.Errorf("Ugyldig verdi: %s", unit)
	}

	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = delimiter

	var total float64
	var count int

	// Looper gjennom hver linje i CSV filen
	for i := 1; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		if i < 2 || i > 16755 {
			// hopper over linjer utenfor rangen
			continue
		}

		if len(record) <= tempColumn {
			return 0, fmt.Errorf("Ugyldig data i filen %s", filename)
		}

		temp, err := strconv.ParseFloat(record[tempColumn], 64)
		if err != nil {
			return 0, err
		}

		total += temp
		count++
	}

	if count == 0 {
		return 0, fmt.Errorf("Ingen tempratur ble funnet i filen %s", filename)
	}

	return total / float64(count), nil
}
