package ica

import (
	"encoding/json"
	"errors"
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
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	return &Service{
		client:    client,
		random:    rand.New(rand.NewSource(time.Now().UnixNano())),
		searchURL: sURL,
		tokenURL:  tURL,
	}
}

func (s *Service) AccessToken() (string, error) {
	req, err := http.NewRequest("GET", s.tokenURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")

	res, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	data := struct {
		AccessToken string `json:"accessToken"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	return data.AccessToken, nil
}

func (s *Service) RecipeCard(token, path string, excludedIDs ...string) (RecipeCard, error) {
	fmt.Print(".")
	url := fmt.Sprintf("%s?url=%s/huvudratt/&onlyEnabled=true&sortOption=rating&take=48", s.searchURL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RecipeCard{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := s.client.Do(req)
	if err != nil {
		return RecipeCard{}, err
	}
	defer res.Body.Close()

	data := struct {
		PageDTO struct {
			RecipeCards []RecipeCard `json:"recipeCards"`
		} `json:"pageDto"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return RecipeCard{}, err
	}

	filteredRecipeCards := filteredRecipeCards(data.PageDTO.RecipeCards, excludedIDs...)

	if len(filteredRecipeCards) == 0 {
		return RecipeCard{}, errors.New("no recipe cards found")
	}

	return filteredRecipeCards[s.random.Intn(len(filteredRecipeCards))], nil
}

func filteredRecipeCards(recipeCards []RecipeCard, excludeIDs ...string) []RecipeCard {
	excludeMap := make(map[string]bool)
	for _, id := range excludeIDs {
		excludeMap[id] = true
	}

	filteredRecipeCards := []RecipeCard{}
	for _, recipeCard := range recipeCards {
		if !excludeMap[strconv.Itoa(recipeCard.ID)] {
			filteredRecipeCards = append(filteredRecipeCards, recipeCard)
		}
	}

	return filteredRecipeCards
}
