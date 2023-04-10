package yr

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	/*"github.com/NorwegianKiwi-glitch/funtemps/conv"*/)

// konverterer temperatur fra Celsius til Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

/*
func Convert() error {
	// Sjekker at output fil ikke eksisterer fra før av
	if _, err := os.Stat("yr/kjevik-temp-fahr-20220318-20230318.csv"); !os.IsNotExist(err) {
		// Output file already exists, prompt user to regenerate
		var regenerate string
		fmt.Print("Output filen eksisterer fra før av, ønsker du å regenerere filen? (y/n): ")
		fmt.Scanln(&regenerate)
		if regenerate != "y" && regenerate != "Y" {
			fmt.Println("Avbrt uten å generere ny fil.")
			return nil
		}
	}

	// Åpner input fil
	inputFile, err := os.Open("yr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Feil input fil:", err)
	}
	defer inputFile.Close()

	// Lager ny scanner for å lese input fil
	inputScanner := bufio.NewScanner(inputFile)

	// lager en ny CSV writer for å skrive output fil
	outputFile, err := os.Create("yr/kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		fmt.Println("Feil generering av output fil:", err)
	}
	defer outputFile.Close()

	outputWriter := csv.NewWriter(outputFile)
	defer outputWriter.Flush()

	// Skriver ut første linje i input fil
	if inputScanner.Scan() {
		firstLine := inputScanner.Text()
		if err = outputWriter.Write(strings.Split(firstLine, ";")); err != nil {
			fmt.Println("Feil under skriving av første linje:", err)
		}
	}

	// Looper gjennom hver linje i input fil
	lineNum := 1 // Inisialiser teller for linjenummer
	for inputScanner.Scan() {
		lineNum++ // Øk linjenummer teller

		// Hopp over linje 2 og 16756
		if lineNum == 2 || lineNum == 16756 {
			continue
		}

		// Splitter linjen i felt
		fields := strings.Split(inputScanner.Text(), ";")

		// Sjekker at feltet har minst 4 elementer
		if len(fields) < 4 {
			fmt.Println("Feil: Ugyldig input format.")
			continue
		}

		// henter ut siste siffer i det 4 feltet
		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			fmt.Println("Feil ved analyse av temperatur:", err)
			continue
		}
		lastDigit := temperature - float64(int(temperature/10))*10

		// Konverterer Celsius til Fahrenheit
		// Ekskulderer linjene 2 og 16756
		if inputScanner.Text() == "" || inputScanner.Text() == "Observasjoner gitt i Norsk tid." || inputScanner.Text() == ";;;,celsius,fahrenheit;," || inputScanner.Text() == "2022-03-18;00:00;2022-03-18;00:00;1,2;4,4;39,9;103,8;" || inputScanner.Text() == "2023-03-18;00:00;2023-03-18;00:00;1,2;7,5;45,5;117,9;" || strings.HasPrefix(inputScanner.Text(), ";Observasjoner") || strings.HasPrefix(inputScanner.Text(), ";;;") {
			continue
		}

		// Konverterer Celsius til Fahrenheit
		fahrenheit := celsiusToFahrenheit(lastDigit)

		// Erstatter det 4 feltet med den konverterte temperaturen
		temperatureString := strconv.FormatFloat(fahrenheit, 'f', 2, 64)
		temperatureParts := strings.Split(temperatureString, ".")
		fields[3] = temperatureParts[0] + "." + string(temperatureParts[1][0])

		// Skriver den oppdaterte linjen til output CSV filen
		err = outputWriter.Write(fields)
		if err != nil {
			fmt.Println("Feil skriving av output linje til fil:", err)
			continue
		}
	}

	dataText := "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Simon Helgen,,,"
	err = outputWriter.Write([]string{dataText})
	if err != nil {
		fmt.Println("Feil under skriving av data tekst til output fil:", err)
	}

	return nil
}
*/

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
		if len(fields) < 4 {
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
		lastDigit := temperature - float64(int(temperature/10))*10

		// Convert Celsius to Fahrenheit
		/*fahrenheit := conv.CelsiusToFarenheit(lastDigit)*/
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
