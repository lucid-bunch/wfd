package ica

import "fmt"

type RecipeCard struct {
	ID          int    `json:"id"`
	AbsolutURL  string `json:"absoluteUrl"`
	CookingTime string `json:"cookingTime"`
	Difficulty  string `json:"difficulty"`
	Title       string `json:"title"`
}

func (r RecipeCard) String() string {
	return fmt.Sprintf("%s\n\n\tDifficulty: %s, %s\n\tLink: %s", r.Title, r.Difficulty, r.CookingTime, r.AbsolutURL)
}
