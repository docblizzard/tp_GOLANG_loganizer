package cmd

import (
	"fmt"
	"sync"

	"github.com/docblizzard/loganizer/internal/checker"
	"github.com/docblizzard/loganizer/internal/config"
	"github.com/docblizzard/loganizer/internal/reporter"
	"github.com/spf13/cobra"
)

var (
	configPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Ajoute une nouvelle URL à un fichier JSON de configuration.",
	Long:  `La commande 'analyze' permet de vérifier un fichier JSON de log.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configPath == "" {
			fmt.Println("Erreur: config path (--config) obligatoire.")
			return
		}

		targets, err := config.LoadTargetsFromFile(configPath)
		if err != nil {
			fmt.Printf("Erreur lors du chargement du fichier: %v\n", err)
			return
		}

		if len(targets) == 0 {
			fmt.Println("Aucun log à lire")
			return
		}

		var wg sync.WaitGroup
		resultsChan := make(chan config.OutputTarget, len(targets))

		wg.Add(len(targets))
		for _, id := range targets {
			go func(t config.InputTarget) {
				defer wg.Done()
				result := checker.ParseLog(t)
				resultsChan <- result // Envoyer le résultat au channel
			}(id)
		}
		// Cette ligne bloque l'éxecution du main() jusqu'à ce que toutes les goroutines aient appelé wd.Done()
		wg.Wait()
		close(resultsChan)
		var finalReport []config.OutputTarget

		for res := range resultsChan { // Récupérer tous les résultats du channel
			finalReport = append(finalReport, res)

			if err := reporter.ExportResultsToJsonfile(finalReport); err != nil {
				fmt.Printf("Erreur lors de l'exportation des résultats : %v\n", err)
			} else {
				fmt.Printf("Résultats exportés dans le dossier log")
			}
		}
	},
}

func init() {
	// Cette ligne est cruciale : elle "ajoute" la sous-commande `checkCmd` à la commande racine `rootCmd`.
	// C'est ainsi que Cobra sait que 'check' est une commande valide sous 'gowatcher'.
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configPath, "config", "i", "", "Absolute path to JSON file containing logs")

	analyzeCmd.MarkFlagRequired("config")
}
