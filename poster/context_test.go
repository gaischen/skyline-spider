package poster

import (
	"context"
	"testing"
	"time"
)

var runContext, cancelFunc = context.WithCancel(context.Background())

func TestCancelFunc(t *testing.T) {

	ti := time.NewTimer(time.Second * 5)
	select {
	case <-ti.C:
		cancelFunc()
		//case <-:

	}
}

func Sleep(ctx context.Context) chan struct{} {
	time.Sleep(time.Second)
	return nil
}
