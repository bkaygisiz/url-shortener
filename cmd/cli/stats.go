package cli

import (
	"fmt"
	"log"
	"os"

	cmd2 "github.com/bkaygisiz/url-shortener/cmd"
	"github.com/bkaygisiz/url-shortener/internal/repository"
	"github.com/bkaygisiz/url-shortener/internal/services"
	"github.com/spf13/cobra"

	"gorm.io/driver/sqlite" // Driver SQLite pour GORM
	"gorm.io/gorm"
)

// Variable pour stocker la valeur du flag --code
var shortCodeFlag string

// StatsCmd représente la commande 'stats'
var StatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Affiche les statistiques (nombre de clics) pour un lien court.",
	Long: `Cette commande permet de récupérer et d'afficher le nombre total de clics
pour une URL courte spécifique en utilisant son code.

Exemple:
  url-shortener stats --code="xyz123"`,
	Run: func(cmd *cobra.Command, args []string) {
		// Valider que le flag --code a été fourni
		if shortCodeFlag == "" {
			fmt.Fprintf(os.Stderr, "Erreur: Le flag --code est requis\n")
			os.Exit(1)
		}

		// Charger la configuration chargée globalement via cmd.cfg
		cfg := cmd2.Cfg
		if cfg == nil {
			log.Fatalf("FATAL: Configuration non chargée")
		}

		// Initialiser la connexion à la base de données SQLite avec GORM
		db, err := gorm.Open(sqlite.Open(cfg.Database.Name), &gorm.Config{})
		if err != nil {
			log.Fatalf("FATAL: Échec de la connexion à la base de données: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("FATAL: Échec de l'obtention de la base de données SQL sous-jacente: %v", err)
		}

		// S'assurer que la connexion est fermée à la fin de l'exécution de la commande
		defer sqlDB.Close()

		// Initialiser les repositories et services nécessaires NewLinkRepository & NewLinkService
		linkRepo := repository.NewLinkRepository(db)
		clickRepo := repository.NewClickRepository(db)
		linkService := services.NewLinkService(linkRepo, clickRepo)

		// Appeler GetLinkStats pour récupérer le lien et ses statistiques
		// Attention, la fonction retourne 3 valeurs
		link, totalClicks, err := linkService.GetLinkStats(shortCodeFlag)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Fprintf(os.Stderr, "Erreur: Aucun lien trouvé avec le code: %s\n", shortCodeFlag)
			} else {
				fmt.Fprintf(os.Stderr, "Erreur lors de la récupération des statistiques: %v\n", err)
			}
			os.Exit(1)
		}

		fmt.Printf("Statistiques pour le code court: %s\n", link.ShortCode)
		fmt.Printf("URL longue: %s\n", link.LongURL)
		fmt.Printf("Total de clics: %d\n", totalClicks)
	},
}

// init() s'exécute automatiquement lors de l'importation du package.
// Il est utilisé pour définir les flags que cette commande accepte.
func init() {
	// Définir le flag --code pour la commande stats
	StatsCmd.Flags().StringVar(&shortCodeFlag, "code", "", "Code court pour lequel afficher les statistiques")

	// Marquer le flag comme requis
	StatsCmd.MarkFlagRequired("code")

	// Ajouter la commande à RootCmd
	cmd2.RootCmd.AddCommand(StatsCmd)
}
