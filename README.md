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

Se mer på test "en test som sjekker at gjennomsnittempraturen er 8.56"