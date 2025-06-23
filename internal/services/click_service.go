package services

import (
	"fmt"

	"github.com/bkaygisiz/url-shortener/internal/models"
	"github.com/bkaygisiz/url-shortener/internal/repository" // Importe le package repository
)

// ClickService est une structure qui fournit des méthodes pour la logique métier des clics.
// Elle est composée de clickRepo qui est de type ClickRepository
type ClickService struct {
	clickRepo repository.ClickRepository
}

// NewClickService crée et retourne une nouvelle instance de ClickService.
// C'est la fonction recommandée pour obtenir un service, assurant que toutes ses dépendances sont injectées.
func NewClickService(clickRepo repository.ClickRepository) *ClickService {
	return &ClickService{
		clickRepo: clickRepo,
	}
}

// RecordClick enregistre un nouvel événement de clic dans la base de données.
// Cette méthode est appelée par le worker asynchrone.
func (s *ClickService) RecordClick(click *models.Click) error {
	// Appeler le ClickRepository pour créer l'enregistrement de clic.
	// Gérer toute erreur provenant du repository.
	if err := s.clickRepo.CreateClick(click); err != nil {
		return fmt.Errorf("failed to record click: %w", err)
	}
	return nil
}

// GetClicksCountByLinkID récupère le nombre total de clics pour un LinkID donné.
// Cette méthode pourrait être utilisée par le LinkService pour les statistiques, ou directement par l'API stats.
func (s *ClickService) GetClicksCountByLinkID(linkID uint) (int, error) {
	// Appeler le ClickRepository pour compter les clics par LinkID.
	count, err := s.clickRepo.CountClicksByLinkID(linkID)
	if err != nil {
		return 0, fmt.Errorf("failed to get clicks count for link ID %d: %w", linkID, err)
	}
	return count, nil
}
