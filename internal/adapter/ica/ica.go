package ica

import (
	"encoding/json"
	"net/http"
)

type Service struct {
	client *http.Client
}

func NewService() *Service {
	client := http.Client{}
	return &Service{
		client: &client,
	}
}

func (s *Service) AccessToken(url string) string {
	req, err := http.NewRequest("GET", url, nil)
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

func (s *Service) RecipeCards(url, token string) []RecipeCard {
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

	return data.PageDTO.RecipeCards
}
