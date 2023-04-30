package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	tURL := os.Getenv("TOKEN_URL")
	token := getToken(tURL)
	sURL := os.Getenv("SEARCH_URL") + "?url="
	filters := "/huvudratt/barn/"
	seed := 48
	sleep := 3 * time.Second
	params := fmt.Sprintf("&onlyEnabled=true&sortOption=rating&take=%d", seed)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Print("\nWhat's For Dinner v1.0\n")
	fmt.Print("======================\n\n")
	time.Sleep(sleep)
	fmt.Printf(" Mån:\t%s\n\n", getResults(fmt.Sprintf("%s/vegetarisk%s%s", sURL, filters, params), token)[r.Intn(seed)])
	time.Sleep(sleep)
	fmt.Printf(" Tis:\t%s\n\n", getResults(fmt.Sprintf("%s/fisk%s%s", sURL, filters, params), token)[r.Intn(seed)])
	time.Sleep(sleep)
	fmt.Printf(" Ons:\t%s\n\n", getResults(fmt.Sprintf("%s/kyckling%s%s", sURL, filters, params), token)[r.Intn(seed)])
	time.Sleep(sleep)
	fmt.Printf("Tors:\t%s\n\n", getResults(fmt.Sprintf("%s/vegetarisk%s%s", sURL, filters, params), token)[r.Intn(seed)])
	time.Sleep(sleep)
	fmt.Printf(" Fre:\t%s\n\n", getResults(fmt.Sprintf("%s/kott%s%s", sURL, filters, params), token)[r.Intn(seed)])
	time.Sleep(sleep)
	os.Exit(0)
}

func getToken(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data := struct {
		AccessToken string `json:"accessToken"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data.AccessToken
}

type RecipeCard struct {
	AbsolutURL  string `json:"absoluteUrl"`
	CookingTime string `json:"cookingTime"`
	Difficulty  string `json:"difficulty"`
	Title       string `json:"title"`
}

func (r RecipeCard) String() string {
	return fmt.Sprintf("%s\n\n\tNivå: %s, %s\n\tLänk: %s", r.Title, r.Difficulty, r.CookingTime, r.AbsolutURL)
}

func getResults(url, token string) []RecipeCard {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data := struct {
		PageDTO struct {
			RecipeCards []RecipeCard `json:"recipeCards"`
		} `json:"pageDto"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	return data.PageDTO.RecipeCards
}
