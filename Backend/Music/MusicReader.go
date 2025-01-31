package Music

import (
	"fmt"
	"os"
)

var file = "Backend/TestMusic/TestMusic.m4a"

func TgetMusic() []byte {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d", len(data))
	return data
}
