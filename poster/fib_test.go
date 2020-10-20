package poster

import (
	"fmt"
	"reflect"
	"testing"
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
