package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"db/internal/initialization"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	initializer, err := initialization.NewInitializer()
	if err != nil {
		log.Fatal(err)
	}

	db, err := initializer.StartDatabase()
	if err != nil {
		log.Fatal(err)
	}

	logger, _ := zap.NewProduction()
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		fmt.Print("[db] > ")
		query, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				logger.Fatal("connection was closed", zap.Error(err))
			}

			logger.Error("failed to read user query", zap.Error(err))
		}

		result := db.HandleQuery(ctx, query)

		fmt.Println(result)
	}
}
