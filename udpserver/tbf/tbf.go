// this package is Single Token Bucket Filter
package tbf

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

// Number of processing bits per second
const (
	M = 800
)

// InitTokenBucket is init context
func InitTokenBucket() (context.Context, *rate.Limiter) {
	ctx := context.Background()
	// Charge Bucket M[bit] every second
	c := rate.Every(time.Second / M)
	// Create Tocken (M is burst size)
	l := rate.NewLimiter(c, M)

	return ctx, l
}

// TokenBucketFilter is limit
func TokenBucketFilter(ctx context.Context, n int, limit *rate.Limiter) error {
	err := limit.WaitN(ctx, n)
	if err != nil {
		return err
	}

	return nil
}
