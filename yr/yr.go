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

// Konverterer temperatur fra Celsius til Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func Convert() error {
	// Sjekker at output fil ikke eksisterer fra før av
	if _, err := os.Stat("yr/kjevik-temp-fahr-20220318-20230318.csv"); !os.IsNotExist(err) {
		// Output file already exists, prompt user to regenerate
		var regenerate string
		fmt.Print("Output fil eksisterer fra før av, ønsker du å regenerere filen? (y/n): ")
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
	for inputScanner.Scan() {
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
		/*fahrenheit := conv.CelsiusToFarenheit(lastDigit)*/
		fahrenheit := celsiusToFahrenheit(lastDigit)

		// Ersaetter det 4 feltet med den konverterte temperaturen
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
		return 0, fmt.Errorf("Ugyldig tempratur verdi: %s", unit)
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
			// hopper over linjer utenfor intervallet
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
		return 0, fmt.Errorf("Ingen tempratur data funnet i filen %s", filename)
	}

	return total / float64(count), nil
}
