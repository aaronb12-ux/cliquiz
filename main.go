package main

import (

	"encoding/csv" //for file reading
	"fmt" //for printing
	"log" //errors
	"os" //opening the file

)

func openFile() [][]string {

    f, err := os.Open("problems.csv")

	if err != nil { //fails when file is found
		log.Fatalf("Unable to read input file")
	}

	defer f.Close()

	csvReader := csv.NewReader(f) //new reader that reads from 'f'

	records, err := csvReader.ReadAll() //reading all of the contents
										//This stores each line into an array 
										//Each line is its own element where each comma separated value is an index

	if err != nil { //fails when there is no comma on line
		log.Fatal("Error reading file contents")
	}

	return records

}


func getQandA(file_data [][]string ) ([]string, []string) { //creating separate questions an answers array from overall file structure

	questions := []string{}
	answers := []string{}

	for i := 0; i < len(file_data); i++ {
		questions = append(questions, file_data[i][0]) 
		answers = append(answers, file_data[i][1])
	}

	return questions, answers

}

func takeQuiz(questions []string, answers []string) int {
	//display the question
	//scan in an answer
	//check if the answer is equal to the
	var answer string
	var numCorrect int

	for i := 0; i < len(questions); i ++ {
		fmt.Print(questions[i] + " ")
		fmt.Scanf("%s", &answer)

		

		if answer == answers[i] {
			numCorrect = numCorrect + 1
		}	
	}

	return numCorrect
}
func showResults(numberCorrect int, numberQuestions int) {

	if numberCorrect == numberQuestions {
		fmt.Print("100%. Great work!")
	} else{
		fmt.Printf("You got %d out of %d correct", numberCorrect, numberQuestions)
	}
}

func main() {

	file_data := openFile()

	questions, answers := getQandA(file_data)

	num_correct := takeQuiz(questions, answers)

	showResults(num_correct, len(answers))


}




