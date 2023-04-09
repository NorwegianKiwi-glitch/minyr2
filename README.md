# minyr2

Porblemer med **imports github.com/NorwegianKiwi-glitch/Funtemps/conv: import cycle not allowed**
    Fikset det med å lage ett nytt reposetory som med navn "minyr2" istedet for "minyr" og bruker "github.com/NorwegianKiwi-glitch/minyr2" i go.mod isetedet for "github.com/NorwegianKiwi-glitch/funtemps/conv"

Problemer med importering av **github.com/NorwegianKiwi-glitch**
    PS C:\Users\simon\reps\IS-105\go\minyr2> go get github.com/NorwegianKiwi-glitch/Funtemps/conv
    go: github.com/NorwegianKiwi-glitch/Funtemps@v0.0.0-20230330101830-93228ad389d7: parsing go.mod:
        module declares its path as: github.com/NorwegianKiwi-glitch/funtemps
                but was required as: github.com/NorwegianKiwi-glitch/Funtemps
    
    Ting jeg har prøvd
        - Gjøre reposetoriet åpent
        - GOPROXY 
        - lage nytt reposetory
        - Feilsøking og hodeskraing i mange timer

Når jeg bruker funtemps/conv fra github blir output i csv fil slik:
Navn,Stasjon,Tid(norsk normaltid),Lufttemperatur
Kjevik,SN39040,18.03.2022 01:50,38.0
Kjevik,SN39040,18.03.2022 03:20,36.0
Kjevik,SN39040,18.03.2022 03:50,35.0
Kjevik,SN39040,18.03.2022 04:20,37.0
Kjevik,SN39040,18.03.2022 04:50,37.0
Kjevik,SN39040,18.03.2022 05:20,37.0
Kjevik,SN39040,18.03.2022 05:50,37.0
Kjevik,SN39040,18.03.2022 06:20,36.0

Når jeg bruker en lokal funksjon for å konvertere tempraturene blir output slik
Navn,Stasjon,Tid(norsk normaltid),Lufttemperatur
Kjevik,SN39040,18.03.2022 01:50,42.8
Kjevik,SN39040,18.03.2022 03:20,39.2
Kjevik,SN39040,18.03.2022 03:50,37.4
Kjevik,SN39040,18.03.2022 04:20,41.0
Kjevik,SN39040,18.03.2022 04:50,41.0
Kjevik,SN39040,18.03.2022 05:20,41.0
Kjevik,SN39040,18.03.2022 05:50,41.0
Kjevik,SN39040,18.03.2022 06:20,39.2

Nesten lik formel blir brukt. Slik ser formelen ut:
lokal funksjon: 
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}
github funksjon:
func CelsiusToFarenheit(value float64) float64 {
	return value* 1.8 + 32
}

Jeg velger derfor å bruke en lokal funksjon i yr.go fremfor å importere funksjonen fra funtemps/conv på grunn av at den konverterer riktig og er mer nøyaktig


Se mer på: test "en test som sjekker at gjennomsnittempraturen er 8.56"
           test "gitt "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;" ønsker å få (want)"Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort avSTUDENTENS_NAVN", hvor STUDENTENS_NAVN er navn på studenten som leverer besvarelsen"