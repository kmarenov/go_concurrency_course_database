package in_memory

import (
	"errors"

	"go.uber.org/zap"
)

type hashTable interface {
	Set(string, string)
	Get(string) (string, bool)
	Del(string)
}

type Engine struct {
	logger *zap.Logger
	table  hashTable
}

func NewEngine(tableBuilder func() hashTable, logger *zap.Logger) (*Engine, error) {
	if tableBuilder == nil {
		return nil, errors.New("hash table builder is invalid")
	}

	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	return &Engine{
		logger: logger,
		table:  tableBuilder(),
	}, nil
}

func (e *Engine) Set(key, value string) {
	e.table.Set(key, value)
	e.logger.Debug("success set query")
}

func (e *Engine) Get(key string) (string, bool) {
	value, found := e.table.Get(key)
	e.logger.Debug("success get query")
	return value, found
}

func (e *Engine) Del(key string) {
	e.table.Del(key)
	e.logger.Debug("success del query")
}
