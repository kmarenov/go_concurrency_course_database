package initialization

import (
	"go.uber.org/zap"

	"db/internal/database/storage"
	"db/internal/database/storage/engine/in_memory"
)

func CreateEngine(logger *zap.Logger) (storage.Engine, error) {
	return in_memory.NewEngine(in_memory.HashTableBuilder, logger)
}
