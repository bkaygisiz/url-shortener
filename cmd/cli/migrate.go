package cli

import (
	"log"

	cmd2 "github.com/bkaygisiz/url-shortener/cmd"
	"github.com/bkaygisiz/url-shortener/internal/models"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MigrateCmd représente la commande 'migrate'
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Exécute les migrations de base de données.",
	Long: `Cette commande exécute les migrations de base de données pour créer ou mettre à jour les tables.
Elle se connecte à la base de données SQLite configurée et exécute les migrations automatiques de GORM
pour les tables 'links' et 'clicks' basées sur les modèles Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Charger la configuration chargée globalement via cmd.cfg
		cfg := cmd2.Cfg
		if cfg == nil {
			log.Fatalf("FATAL: Configuration non chargée")
		}

		// Initialiser la connexion à la base de données SQLite
		db, err := gorm.Open(sqlite.Open(cfg.Database.Name), &gorm.Config{})
		if err != nil {
			log.Fatalf("FATAL: Échec de la connexion à la base de données: %v", err)
		}

		// S'assurer que la connexion est fermée après migration
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("FATAL: Échec de l'obtention de la base de données SQL sous-jacente: %v", err)
		}
		defer sqlDB.Close()

		// Exécuter les migrations avec db.AutoMigrate()
		err = db.AutoMigrate(&models.Link{}, &models.Click{})
		if err != nil {
			log.Fatalf("FATAL: Échec de la migration: %v", err)
		}

		log.Printf("Migrations exécutées avec succès pour la base de données: %s", cfg.Database.Name)
	},
}

func init() {
	// Ajouter la commande à RootCmd
	cmd2.RootCmd.AddCommand(MigrateCmd)
}
