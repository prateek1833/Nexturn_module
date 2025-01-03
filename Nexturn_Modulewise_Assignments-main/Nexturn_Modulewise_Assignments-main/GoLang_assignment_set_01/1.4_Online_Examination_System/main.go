package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"
)

// Constants for exam configuration
const (
    QUESTION_TIME_LIMIT = 30 * time.Second // 30 seconds per question
    QUIT_COMMAND        = "quit"
    PASS_MARKS_PERCENT  = 60.0
)

// Performance levels
const (
    OUTSTANDING = 90.0
    VERY_GOOD   = 75.0
    ACCEPTABLE  = 60.0
)

// Problem represents a single quiz problem
type Problem struct {
    Statement      string
    Choices        []string
    RightAnswer    int
}

// Examination handles the quiz operations
type Examination struct {
    Problems        []Problem
    TotalScore      int
    QuestionCount   int
    InputReader     *bufio.Scanner
}

// NewExamination initializes a new exam session
func NewExamination() *Examination {
    questions := []Problem{
        {
            Statement: "Which city is the capital of France?",
            Choices: []string{
                "Berlin",
                "Madrid",
                "Paris",
                "Rome",
            },
            RightAnswer: 3,
        },
        {
            Statement: "What is the chemical symbol for water?",
            Choices: []string{
                "CO2",
                "H2O",
                "NaCl",
                "O2",
            },
            RightAnswer: 2,
        },
        {
            Statement: "Solve: 5 × 3 - 4",
            Choices: []string{
                "15",
                "11",
                "13",
                "9",
            },
            RightAnswer: 2,
        },
        {
            Statement: "Who is the author of 'Pride and Prejudice'?",
            Choices: []string{
                "Jane Austen",
                "Charles Dickens",
                "George Eliot",
                "Charlotte Brontë",
            },
            RightAnswer: 1,
        },
        {
            Statement: "Which continent is the Sahara Desert located in?",
            Choices: []string{
                "Asia",
                "Africa",
                "Australia",
                "South America",
            },
            RightAnswer: 2,
        },
    }

    return &Examination{
        Problems:      questions,
        QuestionCount: len(questions),
        InputReader:   bufio.NewScanner(os.Stdin),
    }
}

// showQuestion displays a problem and its choices
func (e *Examination) showQuestion(index int, problem Problem) {
    fmt.Printf("\nQuestion %d/%d:\n", index+1, e.QuestionCount)
    fmt.Println(problem.Statement)
    for i, choice := range problem.Choices {
        fmt.Printf("%d. %s\n", i+1, choice)
    }
    fmt.Printf("\nEnter your answer (1-%d) or '%s' to exit: ", len(problem.Choices), QUIT_COMMAND)
}

// getInput reads and validates user responses
func (e *Examination) getInput() (int, error) {
    e.InputReader.Scan()
    response := strings.TrimSpace(e.InputReader.Text())

    if strings.ToLower(response) == QUIT_COMMAND {
        return -1, nil
    }

    choice, err := strconv.Atoi(response)
    if err != nil {
        return 0, fmt.Errorf("please enter a valid number")
    }

    return choice, nil
}

// determinePerformance evaluates the performance category based on score percentage
func determinePerformance(percentage float64) string {
    switch {
    case percentage >= OUTSTANDING:
        return "Outstanding"
    case percentage >= VERY_GOOD:
        return "Very Good"
    case percentage >= ACCEPTABLE:
        return "Acceptable"
    default:
        return "Needs Improvement"
    }
}

// BeginExam starts the quiz
func (e *Examination) BeginExam() {
    fmt.Println("\nWelcome to the Interactive Quiz Platform!")
    fmt.Printf("You have %v for each question. Total questions: %d\n", QUESTION_TIME_LIMIT, e.QuestionCount)
    fmt.Println("Press Enter to begin...")
    e.InputReader.Scan()

    for i, problem := range e.Problems {
        // Create timer for the question
        questionTimer := time.NewTimer(QUESTION_TIME_LIMIT)
        answerChannel := make(chan int)
        timeoutChannel := make(chan bool)

        // Display the question
        e.showQuestion(i, problem)

        // Goroutine to handle user response
        go func() {
            for {
                answer, err := e.getInput()
                if err != nil {
                    fmt.Printf("Error: %v\n", err)
                    continue
                }
                answerChannel <- answer
                break
            }
        }()

        // Goroutine to handle timeout
        go func() {
            <-questionTimer.C
            timeoutChannel <- true
        }()

        // Await user input or timeout
        select {
        case response := <-answerChannel:
            questionTimer.Stop()
            if response == -1 {
                fmt.Println("\nQuiz terminated by the participant.")
                e.showResults()
                return
            }
            if response < 1 || response > len(problem.Choices) {
                fmt.Println("Invalid selection. No points scored.")
                continue
            }
            if response == problem.RightAnswer {
                fmt.Println("Correct!")
                e.TotalScore++
            } else {
                fmt.Printf("Incorrect. The correct answer was: %d\n", problem.RightAnswer)
            }

        case <-timeoutChannel:
            fmt.Println("\nTime's up! Proceeding to the next question...")
        }
    }

    e.showResults()
}

// showResults displays the overall quiz performance
func (e *Examination) showResults() {
    percentage := (float64(e.TotalScore) / float64(e.QuestionCount)) * 100
    performance := determinePerformance(percentage)

    fmt.Println("\n--- Final Results ---")
    fmt.Printf("Total Questions: %d\n", e.QuestionCount)
    fmt.Printf("Correct Answers: %d\n", e.TotalScore)
    fmt.Printf("Score Percentage: %.2f%%\n", percentage)
    fmt.Printf("Performance: %s\n", performance)

    if percentage >= PASS_MARKS_PERCENT {
        fmt.Println("Congratulations! You passed the quiz!")
    } else {
        fmt.Println("Keep practicing and try again. You'll improve!")
    }
}

func main() {
    exam := NewExamination()
    exam.BeginExam()
}
