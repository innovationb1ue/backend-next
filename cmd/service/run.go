package service

import (
	"context"
	"time"

	"go.uber.org/fx"
	// "github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"

	"github.com/penguin-statistics/backend-next/internal/config"
)

func run(app *fiber.App, config *config.Config, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errChan := make(chan error)

			go func() {
				errChan <- app.Listen(config.Address)
			}()

			// wait for at maximum 100ms for errChan to return anything, else just assume succeeded and return
			// nil for error
			select {
			case err := <-errChan:
				return err
			case <-time.After(100 * time.Millisecond):
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			if config.DevMode {
				return nil
			}
			return app.Shutdown()
		},
	})
}
