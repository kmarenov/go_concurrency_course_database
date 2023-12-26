package initialization

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"db/internal/configuration"
	"db/internal/database"
	"db/internal/database/compute"
	"db/internal/database/storage"
)

type Initializer struct {
	engine storage.Engine
	logger *zap.Logger
}

func NewInitializer(cfg *configuration.Config) (*Initializer, error) {
	if cfg == nil {
		return nil, errors.New("failed to initialize: config is invalid")
	}

	logger, err := CreateLogger(cfg.Logging)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	dbEngine, err := CreateEngine(cfg.Engine, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize engine: %w", err)
	}

	return &Initializer{
		engine: dbEngine,
		logger: logger,
	}, nil
}

func (i *Initializer) StartDatabase(ctx context.Context) (*database.Database, error) {
	compute, err := i.createComputeLayer()
	if err != nil {
		return nil, err
	}

	storage, err := i.createStorageLayer()
	if err != nil {
		return nil, err
	}

	db, err := database.NewDatabase(compute, storage, i.logger)
	if err != nil {
		i.logger.Error("failed to start application", zap.Error(err))
		return nil, err
	}

	return db, nil
}

func (i *Initializer) createComputeLayer() (*compute.Compute, error) {
	queryParser, err := compute.NewParser(i.logger)
	if err != nil {
		i.logger.Error("failed to start application", zap.Error(err))
		return nil, err
	}

	queryAnalyzer, err := compute.NewAnalyzer(i.logger)
	if err != nil {
		i.logger.Error("failed to start application", zap.Error(err))
		return nil, err
	}

	compute, err := compute.NewCompute(queryParser, queryAnalyzer, i.logger)
	if err != nil {
		i.logger.Error("failed to start application", zap.Error(err))
		return nil, err
	}

	return compute, nil
}

func (i *Initializer) createStorageLayer() (*storage.Storage, error) {
	storage, err := storage.NewStorage(i.engine, nil, i.logger)
	if err != nil {
		i.logger.Error("failed to start application", zap.Error(err))
		return nil, err
	}

	return storage, nil
}
