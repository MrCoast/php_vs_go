// full processing procedure: 1:34.84
// only copy each cell as is: 0:11.81
// only reverse columns order: 0:16.77
// only reverse string in each cell: 0:23.63
// only uppercase each cell: 0:19.42
// only removing all digits from values using regexp: 1:06.56

package main
import (
    "bufio"
    "encoding/csv"
    "os"
    "io"
    "log"
    "strings"
    "regexp"
)

func main() {
    fdIn, err := os.Open("in.csv")
    checkError("Cannot open in.csv", err)
    defer fdIn.Close()

    fdOut, err := os.OpenFile("out.csv", os.O_WRONLY|os.O_CREATE, 0644)
    checkError("Cannot create/write out.csv", err)

    reader := csv.NewReader(bufio.NewReader(fdIn))
    writer := csv.NewWriter(fdOut)

    // regexps in golang are very slow :(
    //regexpToRemoveDigits := regexp.MustCompile(`\d`)
    regexpToRemoveDigits := regexp.MustCompilePOSIX(`[0-9]`)

    for {
        csvRowIn, err := reader.Read()
        // stop at EOF
        if err == io.EOF {
            break
        }

        var csvRowOut []string

        for _, valueIn := range csvRowIn {
            valueOut := "(NULL VALUE)"

            if valueIn != "" {
                valueOut = reverseString(valueIn)
                valueOut = regexpToRemoveDigits.ReplaceAllString(valueOut, "")
                valueOut = strings.ToUpper(valueOut)
            }

            //valueOut := ""
            //valueOut = reverseString(valueIn)
            //valueOut = strings.ToUpper(valueIn)
            //valueOut = regexpToRemoveDigits.ReplaceAllString(valueIn, "")

            csvRowOut = append([]string{valueOut}, csvRowOut...) // reverse cells in a row
            //csvRowOut = append(csvRowOut, valueOut) // append cells in direct order
        }

        err = writer.Write(csvRowOut)
        checkError("Cannot write csv row", err)
    }

    writer.Flush()
    fdOut.Close()
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
        os.Exit(1);
    }
}

// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverseString(input string) string {
    // Get Unicode code points. 
    n := 0
    rune := make([]rune, len(input))
    for _, r := range input { 
        rune[n] = r
        n++
    } 
    rune = rune[0:n]
    // Reverse 
    for i := 0; i < n/2; i++ { 
        rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
    } 
    // Convert back to UTF-8. 
    output := string(rune)

    return output
}

