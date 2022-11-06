package helper

import (
	"context"
	"time"
)

func Schedule(ctx context.Context, action func(ctx context.Context) error, delay time.Duration) error {
	for {
		err := action(ctx)
		if err != nil {
			return err
		}
		select {
		case <-time.After(delay):
			//
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
