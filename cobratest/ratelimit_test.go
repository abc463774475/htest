package cobratest

import (
	"github.com/abc463774475/bbtool/n_log"
	"testing"
	"time"

	"go.uber.org/ratelimit"
)

func TestRateLimit(t *testing.T) {
	rl := ratelimit.New(100) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		n_log.Info("%d   %v",i, now.Sub(prev))
		prev = now
	}
}
