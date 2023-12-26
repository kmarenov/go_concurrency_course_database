package compute

import (
	"errors"

	"go.uber.org/zap"
)

const (
	setQueryArgumentsNumber = 2
	getQueryArgumentsNumber = 1
	delQueryArgumentsNumber = 1
)

var (
	errInvalidSymbol    = errors.New("invalid symbol")
	errInvalidCommand   = errors.New("invalid command")
	errInvalidArguments = errors.New("invalid arguments")
)

type Analyzer struct {
	handlers []func(Query) error
	logger   *zap.Logger
}

func NewAnalyzer(logger *zap.Logger) (*Analyzer, error) {
	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	analyser := &Analyzer{
		logger: logger,
	}

	analyser.handlers = []func(Query) error{
		SetCommandID: analyser.analyzeSetQuery,
		GetCommandID: analyser.analyzeGetQuery,
		DelCommandID: analyser.analyzeDelQuery,
	}

	return analyser, nil
}

func (a *Analyzer) AnalyzeQuery(tokens []string) (Query, error) {
	if len(tokens) == 0 {
		a.logger.Debug("invalid query")
		return Query{}, errInvalidCommand
	}

	command := tokens[0]
	commandID := CommandNameToCommandID(command)
	if commandID == UnknownCommandID {
		a.logger.Debug(
			"invalid command",
			zap.String("command", command),
		)
		return Query{}, errInvalidCommand
	}

	query := NewQuery(commandID, tokens[1:])
	handler := a.handlers[commandID]
	if err := handler(query); err != nil {
		return Query{}, err
	}

	a.logger.Debug(
		"query analyzed",
		zap.Any("query", query),
	)

	return query, nil
}

func (a *Analyzer) analyzeSetQuery(query Query) error {
	if len(query.Arguments()) != setQueryArgumentsNumber {
		a.logger.Debug(
			"invalid arguments for set query",
			zap.Any("args", query.Arguments()),
		)
		return errInvalidArguments
	}

	return nil
}

func (a *Analyzer) analyzeGetQuery(query Query) error {
	if len(query.Arguments()) != getQueryArgumentsNumber {
		a.logger.Debug(
			"invalid arguments for get query",
			zap.Any("args", query.Arguments()),
		)
		return errInvalidArguments
	}

	return nil
}

func (a *Analyzer) analyzeDelQuery(query Query) error {
	if len(query.Arguments()) != delQueryArgumentsNumber {
		a.logger.Debug(
			"invalid arguments for del query",
			zap.Any("args", query.Arguments()),
		)
		return errInvalidArguments
	}

	return nil
}
