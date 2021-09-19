package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func capFirstChar(s string) string {
    for index, value := range s {
        return string(unicode.ToUpper(value)) + s[index+1:]
    }
    return ""
}

func main() {
	var example string
    fmt.Println("Masukkan String")
	fmt.Scan(&example)

    _, err := strconv.Atoi(example)
    if err != nil {
		reg, err := regexp.Compile("[^a-zA-Z]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString := reg.ReplaceAllString(example, "")
	
		data := []rune(processedString)
		result := []rune{}
		for i := len(data) - 1; i >= 0; i-- {
			result = append(result, data[i])
		}
		loweredVal := strings.ToLower(string(result))
		toTitle := capFirstChar(loweredVal)
		fmt.Printf(toTitle)
    } else {
        fmt.Printf("Error(Harus terdapat string)")
    }	
}