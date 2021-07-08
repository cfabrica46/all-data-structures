package hashtable

import (
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {

	h := NewHashTable()

	h.Set("cesar", 17)
	h.Set("arturo", 20)

	fmt.Println(h.Get("cesar"))
	fmt.Println(h.Get("arturo"))
	fmt.Println(h.Get("luis"))
	fmt.Println(h)

	h.Delete("cesar")

	fmt.Println(h.Get("cesar"))

}
