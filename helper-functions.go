package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func retrieveQuestions() []Question {

	// external call
	resp, err := http.Get(fmt.Sprintf("%s/questions", baseURL))
	if err != nil {
		log.Fatalf("Failed to fetch questions: %v", err)
	}
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected response status:", resp.Status)
	}

	// deserialize response body into Questions
	var respBody []Question
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		log.Fatal("Failed to decode response body:", err)
	}

	return respBody
}

func getQuestions(cmd *cobra.Command, args []string) {
	questions := retrieveQuestions()

	// output
	for _, q := range questions {
		fmt.Printf("Q-%d: %s\n", q.ID, q.Query)
		for i, ans := range q.Answers {
			fmt.Printf("(%d)%s", i, ans)
			if i < len(q.Answers)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n\n")
	}
}

func postAnswers(cmd *cobra.Command, args []string) {
	// obtain flags
	user, _ := cmd.Flags().GetString("user")
	numbers, _ := cmd.Flags().GetString("answers")

	// validations
	if len(user) == 0 {
		log.Fatalf("User missing.")
	}
	if len(numbers) == 0 {
		log.Fatalf("At least one answer is required.")
	}

	// retrieve questions
	questions := retrieveQuestions()
	if len(numbers) > len(questions) {
		log.Fatalf("Too many answers.")
	}

	// build request body
	values := strings.Split(numbers, ",")
	answerMatrix := AnswerMatrix{}

	for i, val := range values {
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Print("Error parsing value:", err)
			return
		}

		answerMatrix.Answers = append(answerMatrix.Answers, Answer{QuestionID: questions[i].ID, AnswerID: num})
	}
	answerMatrixJSON, err := json.Marshal(answerMatrix)
	if err != nil {
		log.Print("Error converting JSON value:", err)
		return
	}

	// external call
	resp, err := http.Post(fmt.Sprintf("%s/answers?user=%s", baseURL, user), "application/json", bytes.NewBuffer(answerMatrixJSON))

	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusCreated {
		log.Print("Unexpected response status:", resp.Status)
		return
	}

	// deserialize response body into AnswerMatrix
	var respBody AnswerMatrix
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		log.Print("Failed to decode response body:", err)
		return
	}

	// output
	fmt.Printf("Score: %d\n", respBody.Score)
}

func getUserStats(cmd *cobra.Command, args []string) {
	// obtain flags
	user, _ := cmd.Flags().GetString("user")

	// validations
	if len(user) == 0 {
		log.Fatalf("User missing.")
	}

	// external call
	resp, err := http.Get(fmt.Sprintf("%s/user/%s/statistics", baseURL, user))
	if err != nil {
		log.Fatalf("Failed to fetch user statistics: %v", err)
	}
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		log.Print("Unexpected response status:", resp.Status)
		return
	}

	// deserialize response body into Statistic
	var respBody Statistic
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		log.Print("Failed to decode response body:", err)
		return
	}

	// output
	fmt.Printf("You were better than %d%% of all quizzers\n", int(respBody.WorseQuizzersRatio*100))
}
