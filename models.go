package main

type Question struct {
	ID            int      `json:"id"`
	Query         string   `json:"query"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"-"` // do not expose this field
}

type Answer struct {
	QuestionID int `json:"questionId"`
	AnswerID   int `json:"answerId"`
}

type AnswerMatrix struct {
	Score   int      `json:"score"`
	Answers []Answer `json:"answers"`
}

type Statistic struct {
	UserScore          int     `json:"userScore"`
	WorseQuizzersRatio float32 `json:"worseQuizzersRatio"`
	TotalQuizzers      int     `json:"totalQuizzers"`
}
