package main

import (
	"encoding/csv" //for file reading
	"fmt" //for printing
	"log" //errors
	"os" //opening the file

)

func openFile() [][]string {

    f, err := os.Open("problems.csv")

	if err != nil {
		log.Fatalf("Unable to read input file")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Error reading file contents")
	}

	fmt.Printf("Type of records: %T\n", records)

	return records
	

}

func getQandA(file_data [][]string ) ([]string, []string) {

	questions := []int{}
	answers := []int{}

	for i := range len(file_data) {
		
	}

	

}

func main() {

	file_data := openFile()

	fmt.Println(file_data)

	questions, answers := getQandA(file_data)

}




