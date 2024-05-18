package movie

import "fmt"

func init() {
	fmt.Println("Init: movie")
}

func Review(name string, rating float64) {
	fmt.Printf("I received %s and it's rating is %f\n", name, rating)
}
