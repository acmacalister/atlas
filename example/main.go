package main

import "github.com/acmacalister/atlas"

func main() {
	fe := atlas.Frontend{Host: ":8080"}
	rr := atlas.New(2)
	atlas.Run([]atlas.Frontend{fe}, []string{"localhost:8081", "localhost:8082"}, rr)
}
