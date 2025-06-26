package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganizer",
	Short: "Gowatcher est un outil pour vérifier l'accessibilité des URLs.",
	Long:  `Un outil CLI en Go pour analyzer des fichiers de logs en parallèle pour en extraire des informations clés.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
