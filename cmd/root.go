package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lucid-bunch/wfd/internal/adapter/ica"
	"github.com/lucid-bunch/wfd/internal/store"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wfd",
	Short: "wfd is a tool that helps with dinner planning",
	Long:  "What's For Dinner is a tool that helps with dinner planning",
	Run: func(cmd *cobra.Command, args []string) {
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

		fmt.Printf("- generating recipes")
		svc := ica.NewService(os.Getenv("SEARCH_URL"), os.Getenv("TOKEN_URL"))
		token := svc.AccessToken()

		types, err := cmd.Flags().GetStringArray("type")
		if err != nil {
			fmt.Printf("Error reading recipe types: %s\n\n", err)
		}

		recipes := []ica.RecipeCard{}
		for _, t := range types {
			recipes = append(recipes, svc.RecipeCard(token, "/"+t, excludeIDs...))
			delay()
		}

		fmt.Printf("\n- printing recipes\n\n")
		ids := []string{}
		for _, recipe := range recipes {
			fmt.Printf("%s\n\n", recipe)
			if recipe.ID != 0 {
				ids = append(ids, fmt.Sprintf("%d", recipe.ID))
			}
		}
		delay()

		fmt.Printf("- writing to recipe store\n")
		if err := cStore.Write(ids); err != nil {
			fmt.Printf("Error writing to store: %s\n\n", err)
		}
		delay()

		fmt.Printf("- done\n")
	},
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().StringArrayP("type", "t", []string{"vegetarisk", "fisk", "kyckling", "vegetarisk", "kott"}, "type of recipe")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
