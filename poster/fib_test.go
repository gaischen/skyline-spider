package poster

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestFibList(t *testing.T) {
	var a = 1
	var b = 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestPointer(t *testing.T) {
	a := 1
	ap := &a
	t.Log(ap)
	t.Log(reflect.TypeOf(ap))

}

func TestArray(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arr1 := arr[1:5]
	t.Log(arr1)
	t.Log(len(arr))
	i := 1
	change(i)
	t.Log("change i:", i)
}

func change(i int) {
	i = i + 1
}

func TestSelect(t *testing.T) {
	done := make(chan struct{})
	go ticker(done)
	time.Sleep(3 * time.Second)
	done <- struct{}{}
	close(done)
	time.Sleep(3 * time.Second)
}

func ticker(done chan struct{}) {
	t := time.NewTicker(1 * time.Second)
	for range t.C {
		select {
		case <-done:
			fmt.Println("done exit")
			return
		default:
			fmt.Println("tick")
		}
	}
}

func TestContext(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go tickerContext(ctx)
	time.Sleep(3 * time.Second)
	cancelFunc()
	time.Sleep(3 * time.Second)
}

func tickerContext(ctx context.Context) {
	t := time.NewTicker(1 * time.Second)
	for range t.C {
		select {
		case <-ctx.Done():
			fmt.Println("done...")
			return
		default:
			fmt.Println("tick...")
		}
	}
}
