package main

import (
	"syncenv/functions"
	"flag"
	"fmt"
)

func main() {
	fileName := flag.String("file", ".env.development", "file to save secrets")
	secretName := flag.String("secret", "dev/exposureIQ/fe", "path to AWS secret")

	flag.Parse()

	fmt.Println("Fetching secrets from AWS")
	jsonString := functions.GetSecret(*secretName)
	jsonResult := functions.JsonParse(jsonString)

	fmt.Println("Writing secrets to file:", *fileName)
	functions.WriteFile(*fileName, jsonResult)

	fmt.Println("Successfully wrote secrets to file.")
}
