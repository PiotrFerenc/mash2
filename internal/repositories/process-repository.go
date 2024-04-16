package repositories

import (
	types "github.com/PiotrFerenc/mash2/internal/types"
)

type ProcessRepository interface {
	UpdateStatus(pipeline types.Process)
	Save(pipeline types.Process)
}
