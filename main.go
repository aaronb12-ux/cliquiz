package main

import (
	"encoding/csv" //for file reading
	"flag"
	"fmt" //for printing
	"log" //errors
	"os"  //opening the file
	"time"
	"aaron.com/cliquiz/apple"
)

func openFile() [][]string {

	f, err := os.Open("problems.csv")

	if err != nil { //if there is an error opening the file
		log.Fatalf("Unable to read input file")
	}

	defer f.Close()

	csvReader := csv.NewReader(f) //new reader that reads from 'f'

	records, err := csvReader.ReadAll() //reading all of the contents
	//This stores each line into an array
	//Each line is its own element where each comma separated value is an index

	if err != nil { //fails when there is no comma on line
		log.Fatal("Error reading file contents. Make sure the file is formatted correctly is CSV.")
	}

	return records
}

func getQandA(file_data [][]string) ([]string, []string) { //creating separate questions an answers array from overall file structure

	questions := []string{}
	answers := []string{}

	for i := 0; i < len(file_data); i++ {
		questions = append(questions, file_data[i][0])

		answers = append(answers, file_data[i][1])
	}

	return questions, answers

}

func takeQuiz(questions []string, answers []string, timer *time.Timer) int {

	var numCorrect int

	for i := 0; i < len(questions); i++ {

		fmt.Printf("Problem #%d: %s = ", i+1, questions[i]) 
		answerCh := make(chan string)

		go func() {
			var answer string //scan for an answer and put it in the answer string. To get the answer back and know when its coming to us, we make a channel
			fmt.Scanf("%s", &answer)
			answerCh <- answer //when we get an answer, send it to the answer channel
		}()

		select {
		case <-timer.C: //if we get a message from the channel of the timer
			fmt.Println("\nran out of time!")
			return numCorrect

		case answer := <-answerCh: //if we can an answer from the answerChannel
			if answer == answers[i] {
				numCorrect = numCorrect + 1
			}
		}
	}

	return numCorrect
}

func showResults(numberCorrect int, numberQuestions int) {

	if numberCorrect == numberQuestions {
		fmt.Print("100%. Great work!")
	} else {
		fmt.Printf("You got %d out of %d correct", numberCorrect, numberQuestions)
	}
}

func startQuiz() bool {
	var begin string
	for {
		fmt.Scanf("%s", &begin)
		if begin == "start" {
			return true
		} else {
			fmt.Println("Enter 'start' to begin.")
		}
	}
}

func main() {

	apple.Apple("Sour")
	
	var timeLimit int

	file_data := openFile()

	flag.IntVar(&timeLimit, "timelimit", 5, "the time limit for the quiz in seconds")
	flag.Parse()

	fmt.Printf("You have %d seconds to finish the quiz. Type 'start' and press enter to begin the quiz: ", timeLimit)

	questions, answers := getQandA(file_data)

	start := startQuiz()

	if start {
		timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
		num_correct := takeQuiz(questions, answers, timer)
		showResults(num_correct, len(answers))
	}
}

