package initialization

import (
	"fmt"

	"go.uber.org/zap"

	"db/internal/database"
	"db/internal/database/compute"
	"db/internal/database/storage"
)

type Initializer struct {
	engine storage.Engine
	logger *zap.Logger
}

func NewInitializer() (*Initializer, error) {
	logger, err := CreateLogger()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	dbEngine, err := CreateEngine(logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize engine: %w", err)
	}

	return &Initializer{
		engine: dbEngine,
		logger: logger,
	}, nil
}

func (i *Initializer) StartDatabase() (*database.Database, error) {
	computeLayer, err := i.createComputeLayer()
	if err != nil {
		return nil, err
	}

	storageLayer, err := i.createStorageLayer()
	if err != nil {
		return nil, err
	}

	db, err := database.NewDatabase(computeLayer, storageLayer, i.logger)
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

	queryCompute, err := compute.NewCompute(queryParser, queryAnalyzer, i.logger)
	if err != nil {
		i.logger.Error("failed to start application", zap.Error(err))
		return nil, err
	}

	return queryCompute, nil
}

func (i *Initializer) createStorageLayer() (*storage.Storage, error) {
	s, err := storage.NewStorage(i.engine, i.logger)
	if err != nil {
		i.logger.Error("failed to start application", zap.Error(err))
		return nil, err
	}

	return s, nil
}
