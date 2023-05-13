package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/lucid-bunch/wfd/internal/adapter/ica"
	"github.com/lucid-bunch/wfd/internal/store"
	"github.com/spf13/cobra"
)

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Generate a weekly dinner plan",
	Long:  `Each day of the week will randomly be assigned a recipe from the configured category`,
	Run: func(cmd *cobra.Command, args []string) {
		defaultRecipe := ica.RecipeCard{Title: "Fil, flingor, macka och ägg", Difficulty: "Enkel", CookingTime: "Under 15 min", AbsolutURL: "N/A"}
		cooldown := 4

		path := os.Getenv("STORE_PATH")
		st := store.New(path)

		fmt.Printf("- reading from recipe store\n")
		data, err := st.Read()
		if err != nil {
			fmt.Printf("Error reading from store: %s\n\n", err)
		}
		time.Sleep(2 * time.Second)

		fmt.Printf("- applying %d week recipe cooldown\n", cooldown)
		var excludeIDs []string
		switch {
		case len(data) > cooldown:
			for i := len(data) - cooldown; i < len(data); i++ {
				excludeIDs = append(excludeIDs, data[i]...)
			}
		case len(data) <= cooldown:
			for _, row := range data {
				excludeIDs = append(excludeIDs, row...)
			}
		default:
			excludeIDs = []string{}
		}
		time.Sleep(2 * time.Second)

		fmt.Printf("- generating weekly dinner plans")
		svc := ica.NewService(os.Getenv("SEARCH_URL"), os.Getenv("TOKEN_URL"))
		token := svc.AccessToken()

		week := map[string]ica.RecipeCard{
			"Mon": defaultRecipe,
			"Tue": svc.RecipeCard(token, "/vegetarisk", excludeIDs...),
			"Wed": svc.RecipeCard(token, "/fisk", excludeIDs...),
			"Thu": defaultRecipe,
			"Fri": svc.RecipeCard(token, "/kyckling", excludeIDs...),
			"Sat": svc.RecipeCard(token, "/vegetarisk", excludeIDs...),
			"Sun": svc.RecipeCard(token, "/kott", excludeIDs...),
		}

		ids := []string{}
		for _, recipe := range week {
			if recipe.ID != 0 {
				ids = append(ids, fmt.Sprintf("%d", recipe.ID))
			}
		}

		fmt.Printf("\n- printing weekly dinner plans\n\n")
		fmt.Printf("Mon:\t%s\n\n", week["Mon"])
		fmt.Printf("Tue:\t%s\n\n", week["Tue"])
		fmt.Printf("Wed:\t%s\n\n", week["Wed"])
		fmt.Printf("Thu:\t%s\n\n", week["Thu"])
		fmt.Printf("Fri:\t%s\n\n", week["Fri"])
		fmt.Printf("Sat:\t%s\n\n", week["Sat"])
		fmt.Printf("Sun:\t%s\n\n", week["Sun"])

		fmt.Printf("- writing to recipe store\n")
		if err := st.Write(ids); err != nil {
			fmt.Printf("Error writing to store: %s\n\n", err)
		}
		fmt.Printf("- done\n")
	},
}
