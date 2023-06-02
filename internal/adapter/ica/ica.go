package ica

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Service struct {
	client    *http.Client
	random    *rand.Rand
	searchURL string
	tokenURL  string
}

func NewService(sURL, tURL string) *Service {
	client := http.Client{}
	return &Service{
		client:    &client,
		random:    rand.New(rand.NewSource(time.Now().UnixNano())),
		searchURL: sURL,
		tokenURL:  tURL,
	}
}

func (s *Service) AccessToken() string {
	req, err := http.NewRequest("GET", s.tokenURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")

	res, err := s.client.Do(req)
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

func (s *Service) RecipeCard(token, path string, excludedIDs ...string) RecipeCard {
	fmt.Print(".")
	url := fmt.Sprintf("%s?url=%s/huvudratt/barn/&onlyEnabled=true&sortOption=rating&take=48", s.searchURL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := s.client.Do(req)
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

	filteredRecipeCards := filteredRecipeCards(data.PageDTO.RecipeCards, excludedIDs...)

	return filteredRecipeCards[s.random.Intn(len(filteredRecipeCards))]
}

func filteredRecipeCards(recipeCards []RecipeCard, excludeIDs ...string) []RecipeCard {
	filteredRecipeCards := []RecipeCard{}
	for _, recipeCard := range recipeCards {
		exclude := false
		for _, id := range excludeIDs {
			if strconv.Itoa(recipeCard.ID) == id {
				exclude = true
				break
			}
		}
		if !exclude {
			filteredRecipeCards = append(filteredRecipeCards, recipeCard)
		}
	}

	return filteredRecipeCards
}
