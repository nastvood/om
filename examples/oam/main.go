package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/nastvood/om/conf"
	"github.com/nastvood/om/oam"
	"golang.org/x/exp/rand"
)

func main() {
	capacity := 10

	m := oam.New[int, string](conf.WithCapacity(capacity))

	values := rand.Perm(capacity)
	for _, v := range values {
		m.Add(v, strconv.Itoa(v))
	}

	var i int

	iter := m.Iterator()
	for v, ok := iter.Next(); ok; v, ok = iter.Next() {
		fmt.Println(v)

		if values[i] != v {
			log.Panicf("%d != %d", values[i], v)
		}

		i++
	}
}
