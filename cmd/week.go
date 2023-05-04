package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/lucid-bunch/wfd/internal/adapter/ica"
	"github.com/spf13/cobra"
)

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Generate a weekly dinner plan",
	Long:  `Each day of the week will randomly be assigned a recipe from the configured category`,
	Run: func(cmd *cobra.Command, args []string) {
		svc := ica.NewService()
		tURL := os.Getenv("TOKEN_URL")
		token := svc.AccessToken(tURL)
		sURL := os.Getenv("SEARCH_URL") + "?url="
		filters := "/huvudratt/barn/"
		seed := 48
		sleep := 3 * time.Second
		params := fmt.Sprintf("&onlyEnabled=true&sortOption=rating&take=%d", seed)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		simple := ica.RecipeCard{Title: "Fil, flingor, macka och ägg", Difficulty: "Enkel", CookingTime: "Under 15 min", AbsolutURL: "N/A"}
		time.Sleep(sleep)
		fmt.Printf("Mon:\t%s\n\n", simple)
		time.Sleep(sleep)
		fmt.Printf("Tue:\t%s\n\n", svc.RecipeCards(fmt.Sprintf("%s/vegetarisk%s%s", sURL, filters, params), token)[r.Intn(seed)])
		time.Sleep(sleep)
		fmt.Printf("Wed:\t%s\n\n", svc.RecipeCards(fmt.Sprintf("%s/fisk%s%s", sURL, filters, params), token)[r.Intn(seed)])
		time.Sleep(sleep)
		fmt.Printf("Thu:\t%s\n\n", simple)
		time.Sleep(sleep)
		fmt.Printf("Fri:\t%s\n\n", svc.RecipeCards(fmt.Sprintf("%s/kyckling%s%s", sURL, filters, params), token)[r.Intn(seed)])
		time.Sleep(sleep)
		fmt.Printf("Sat:\t%s\n\n", svc.RecipeCards(fmt.Sprintf("%s/vegetarisk%s%s", sURL, filters, params), token)[r.Intn(seed)])
		time.Sleep(sleep)
		fmt.Printf("Sun:\t%s\n\n", svc.RecipeCards(fmt.Sprintf("%s/kott%s%s", sURL, filters, params), token)[r.Intn(seed)])
		time.Sleep(sleep)
	},
}
