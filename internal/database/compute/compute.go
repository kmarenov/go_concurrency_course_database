package compute

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

type parser interface {
	ParseQuery(string) ([]string, error)
}

type analyzer interface {
	AnalyzeQuery([]string) (Query, error)
}

type Compute struct {
	parser   parser
	analyzer analyzer
	logger   *zap.Logger
}

func NewCompute(parser parser, analyzer analyzer, logger *zap.Logger) (*Compute, error) {
	if parser == nil {
		return nil, errors.New("query parser is invalid")
	}

	if parser == nil {
		return nil, errors.New("query analyzer is invalid")
	}

	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	return &Compute{
		parser:   parser,
		analyzer: analyzer,
		logger:   logger,
	}, nil
}

func (d *Compute) HandleQuery(ctx context.Context, queryStr string) (Query, error) {
	if ctx.Err() != nil {
		d.logger.Debug("query canceled")
		return Query{}, ctx.Err()
	}

	tokens, err := d.parser.ParseQuery(queryStr)
	if err != nil {
		return Query{}, err
	}

	query, err := d.analyzer.AnalyzeQuery(tokens)
	if err != nil {
		return Query{}, err
	}

	return query, nil
}
