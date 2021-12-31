package qa

import (
	"math/rand"
	"time"
)

func AnswerQuestion(question string) string {
	//Create a seed to selecta random slice from a
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return answers[r1.Intn(6)]
}
