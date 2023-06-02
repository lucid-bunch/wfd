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
		delay := func() { time.Sleep(1 * time.Second) }

		cPath := os.Getenv("COOLDOWN_PATH")
		cStore := store.New(cPath)

		fmt.Printf("- reading from recipe cooldown store\n")
		cData, err := cStore.Read()
		if err != nil {
			fmt.Printf("Error reading from recipe cooldown store: %s\n\n", err)
		}
		delay()

		fmt.Printf("- applying %d week recipe cooldown\n", cooldown)
		var excludeIDs []string
		switch {
		case len(cData) > cooldown:
			for i := len(cData) - cooldown; i < len(cData); i++ {
				excludeIDs = append(excludeIDs, cData[i]...)
			}
		case len(cData) <= cooldown:
			for _, row := range cData {
				excludeIDs = append(excludeIDs, row...)
			}
		default:
			excludeIDs = []string{}
		}
		delay()

		bPath := os.Getenv("BLOCK_PATH")
		bStore := store.New(bPath)

		fmt.Printf("- reading from recipe block store\n")
		bData, err := bStore.Read()
		if err != nil {
			fmt.Printf("Error reading from recipe block store: %s\n\n", err)
		}
		delay()

		fmt.Printf("- applying blocked recipe filter\n")
		for _, row := range bData {
			excludeIDs = append(excludeIDs, row...)
		}
		delay()

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
		delay()

		fmt.Printf("\n- printing weekly dinner plans\n\n")
		fmt.Printf("Mon:\t%s\n\n", week["Mon"])
		fmt.Printf("Tue:\t%s\n\n", week["Tue"])
		fmt.Printf("Wed:\t%s\n\n", week["Wed"])
		fmt.Printf("Thu:\t%s\n\n", week["Thu"])
		fmt.Printf("Fri:\t%s\n\n", week["Fri"])
		fmt.Printf("Sat:\t%s\n\n", week["Sat"])
		fmt.Printf("Sun:\t%s\n\n", week["Sun"])
		delay()

		fmt.Printf("- writing to recipe store\n")
		if err := cStore.Write(ids); err != nil {
			fmt.Printf("Error writing to store: %s\n\n", err)
		}
		delay()

		fmt.Printf("- done\n")
	},
}
