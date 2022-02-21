package entity

import (
	"testing"
)

func TestQuestion(t *testing.T) {
	q := Question{
		Id: 1,
		Text: "O que prefere?",
		Answer: ["Comer", "Dormir", "Estudar"],
		ParentId: 0,
	}

}