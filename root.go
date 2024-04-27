package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz-client-cli",
	Short: "CLI Client interacting with Quiz server",
	Long: `Quiz-Client-CLI interacts with Quiz-Server.
	This client can retrieve available questions, post the given answers and also compares the user's score with other participants.`,
}

func execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd, submitAnswersCmd, getUserStatisticsCmd)

	getUserStatisticsCmd.Flags().StringP("user", "u", "", "User id for whom to get statistics")
	submitAnswersCmd.Flags().StringP("user", "u", "", "User id for whom to submit answers")
	submitAnswersCmd.Flags().StringP("answers", "a", "", "Comma-separated list of answers in order of question id")
}

var getQuestionsCmd = &cobra.Command{
	Use:   "get-questions",
	Short: "Retrieve all questions",
	Long:  `Retrieve all questions with their respective multiple choice answers`,
	Run:   getQuestions,
}

var submitAnswersCmd = &cobra.Command{
	Use:   "submit-answers",
	Short: "Submit answers",
	Run:   postAnswers,
}

var getUserStatisticsCmd = &cobra.Command{
	Use:   "get-user-statistics",
	Short: "Get user statistics",
	Run:   getUserStats,
}
