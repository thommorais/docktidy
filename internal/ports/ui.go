package ports

import "github.com/thommorais/docktidy/internal/domain"

// UIService defines operations for the user interface
type UIService interface {
	// Run starts the TUI application
	Run() error

	// DisplayResources shows resources in the UI
	DisplayResources(resources []domain.Resource)

	// DisplayPruneCandidates shows pruning candidates
	DisplayPruneCandidates(candidates []domain.PruneCandidate)

	// ConfirmPrune asks user to confirm pruning operation
	ConfirmPrune(candidates []domain.PruneCandidate) (bool, error)

	// DisplayResult shows the result of a pruning operation
	DisplayResult(result domain.PruneResult)
}
