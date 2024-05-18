package main

import (
	"fmt"

	"github.com/nonybueno/cinema/movie"
	"github.com/nonybueno/cinema/ticket"
)

func init() {
	fmt.Println("Init: main")
}

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)

	// v := "72"
	// s, err := strconv.Atoi(v)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(s)

	// i := 10
	// n := strconv.Itoa(i)
	// fmt.Println(n)

	// fmt.Printf("type:%T, val: %d\n", s, s)

	// var i int = 256
	// fmt.Printf("type:%T, val: %d\n", i, i)

	// var f float64 = float64(i)
	// fmt.Printf("type:%T, val: %f\n", f, f)

	// var u uint8 = uint8(f)
	// fmt.Printf("type:%T, val: %d\n", u, u)

	// s := "If it looks like a duck swims like a duck and quacks like a duck then it probably is a duck"
	// w := WordCount(s)
	// fmt.Printf("%#v\n", w)

	// m := map[string]int{}
	// fmt.Println(m)

	// mm := make(map[string]int)
	// fmt.Println(mm)

	// m["Answer"] = 42
	// fmt.Printf("The value: %#v\n", m)

	// v := m["Answer"]
	// fmt.Println(v)

	// m["Nong"] = 15
	// fmt.Println(m)

	// m["Nong"] = 20
	// fmt.Println(m)

	// delete(m, "Answer")
	// fmt.Printf("The value: %#v\n", m)

	// n := m["Answer"]
	// fmt.Println(n)

	// r, err := Divide(1, 2)
	// if err != nil {
	// 	fmt.Println("handler error:", err)
	// 	return
	// }
	// fmt.Println(r)

	// eg := &movie{
	// 	title:       "Avengers: Endgame",
	// 	year:        2019,
	// 	rating:      8.4,
	// 	votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
	// 	genres:      []string{"Action", "Drama"},
	// 	isSuperHero: true,
	// }
	// vote(eg, 8)
	// fmt.Printf("%#v\n", eg.votes)

	// v := course{}
	// summary(v)
	// show(v)

	// v := "gopher"
	// sale(v)

	// var v any
	// v = 3.2
	// show(v)

	// fmt.Printf("%T %#v\n", v, v)
	// v = "Hello"
	// fmt.Printf("%T %#v\n", v, v)

	// eg := &movie{"Avengers: Endgame", 2019, 8.4, []float64{7, 8, 9, 10, 6, 9, 9, 10, 8}, []string{"Action", "Drama"}, true}

	// eg.addVote(8)
	// fmt.Println("votes:", eg.votes)

	// c := &course{"Basic Go", "Anuchito", 9999}
	// d := c.discount(599)

	// fmt.Println("discount price:", d)
	// fmt.Println("price:", c.price)

	// var price int = 9999
	// var addr *int = &price

	// fmt.Println(price, &price)
	// fmt.Println(addr)

	// *addr = 9400 //write
	// fmt.Println(price)

	// v := *addr //read
	// fmt.Println(v)

	// fmt.Printf("%T\n", addr)

	// ae := movie{
	// 	title:       "Avengers: Endgame",
	// 	year:        2019,
	// 	rating:      8.4,
	// 	genres:      []string{"Action", "Drama"},
	// 	isSuperHero: true,
	// }

	// ae.info()

	// c1 := course{"Basic Go", "Anuchito", 9999}
	// fmt.Printf("%#v\n", c1)

	// d := c1.discount(399)
	// fmt.Println("discount price:", d)
	// movie1 := movie{title: "Avengers: Endgame", year: 2019, rating: 8.4, isSuperHero: true}
	// movie1.genres = append(movie1.genres, "Action", "Drama")
	// movie2 := movie{title: "Avengers: Infinity War", year: 2018, rating: 8.4, isSuperHero: true}
	// movie2.genres = append(movie2.genres, "Action", "Sci-Fi")

	// var mvs []movie

	// mvs = append(mvs, movie1, movie2)

	// for i := 0; i < len(mvs); i++ {
	// 	fmt.Printf("main.movie%#v\n", mvs[i])
	// }

	// c := course{
	// 	name:       "Basic Go",
	// 	instructor: "Anuchit0",
	// 	price:      99999,
	// }

	// c2 := course{"Basic GO", "Anuchito", 9999}
	// c3 := course{instructor: "Anuchito"}

	// fmt.Printf("c2: %#v\n", c2)
	// fmt.Printf("c2: %#v\n", c3)

	// n := c.name
	// fmt.Println("name:", n)

	// c.instructor = "Nong"

	// fmt.Println("name:", c.name)
	// fmt.Println("name:", c.instructor)
	// fmt.Println("name:", c.price)
	// skills := [3]string{"js", "golang", "python"}
	// sum := 1
	// for sum < 5 {
	// 	sum += sum
	// 	fmt.Println("sum:", sum)
	// }
	// fmt.Println(sum)
	// for i := 0; i < len(skills); i++ {
	// 	fmt.Println(skills[i])
	// }

	// for _, value := range skills {
	// 	fmt.Println("value:", value)
	// }

	// genres := [3]string{"Action", "Adventure", "Fantacy"}
	// fmt.Printf("before for loop: %#v\n", genres)

	// for i := 0; i < len(genres); i++ {
	// 	genres[i] = "Movie: " + genres[i]
	// 	// fmt.Println("genres[0]:", genres[i])
	// }

	// fmt.Printf("after for loop: %#v\n", genres)

	// for i, val := range genres {
	// 	fmt.Printf("genres[%d]: %s\n", i, val)
	// }

	// skills := []string{"js", "go", "python"}
	// fmt.Printf("%T => %#v\n", skills, skills)

	// skills = append(skills, "ruby", "java", "erlang")
	// fmt.Printf("length: %d --val: %#v\n", len(skills), skills)

	// show(skills)
	// show(skills[0:2])
	// l := len(skills)
	// show(skills[0:l])
	// show(skills[0:])
	// show(skills[:l])
	// show(skills[:])

	// var ss []int
	// fmt.Printf("%#v\n", ss)

	// ss := make([]int, 3)
	// fmt.Printf("%#v\n", ss)

	// s := skills[0]
	// fmt.Printf("%#v\n", s)

	// for i := 0; i < len(skills); i++ {
	// 	fmt.Printf("for => %#v\n", skills[i])
	// }

	// for _, val := range skills {
	// 	fmt.Printf("for-range => %#v\n", val)
	// }

	// xs := []float64{4, 5, 7, 8, 3, 8, 0}
	// ys := []float64{7, 2, 10, 9, 7}

	// var vote []float64

	// for i := 0; i < len(xs); i++ {
	// 	vote = append(vote, xs[i])
	// }
	// for i := 0; i < len(ys); i++ {
	// 	vote = append(vote, ys[i])
	// }
	// vote := append(xs, ys...)

	// fmt.Printf("vote 5-8: %v", vote[5:9])
}

// func WordCount(s string) map[string]int {
// 	arr := strings.Fields(s)
// 	r := map[string]int{}
// 	for i := 0; i < len(arr); i++ {
// 		_, ok := r[arr[i]]
// 		if ok {
// 			r[arr[i]] += 1
// 		} else {
// 			r[arr[i]] = 1
// 		}
// 	}
// 	return r
// }

// type error interface {
// 	Error() string
// }

// var myErr = errors.New("my custom error message")

// type MyError struct{}

// func (e MyError) Error() string {
// 	return "MyError"
// }

// func Divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		// err := MyError{}
// 		// err := fmt.Errorf("can't divide '%f' by zero", a)
// 		return 0, myErr
// 	}
// 	r := a / b
// 	return r, nil
// }

// func vote(v voter, rating float64) {
// 	v.addVote(rating)
// }

// type movie struct {
// 	title       string
// 	year        int
// 	rating      float64
// 	votes       []float64
// 	genres      []string
// 	isSuperHero bool
// }

// func (m *movie) addVote(rating float64) {
// 	m.votes = append(m.votes, rating)
// }

// type voter interface {
// 	addVote(rating float64)
// }

// type promotion interface {
// 	discount() int
// }

// type infoer interface {
// 	info()
// }

// type presenter interface {
// 	promotion
// 	infoer
// }

// type course struct{}

// func summary(val presenter) {
// 	fmt.Printf("sale: %#v\n", val.discount())
// 	val.info()
// }

// func show(val infoer) {
// 	val.info()
// }

// func (c course) discount() int {
// 	return 599
// }

// func (c course) info() {
// 	fmt.Println("info:", c)
// }

// type movie struct {
// 	title       string
// 	year        int
// 	rating      float64
// 	votes       []float64
// 	genres      []string
// 	isSuperHero bool
// }

//	func (m *movie) addVote(rating float64) {
//		m.votes = append(m.votes, rating)
//	}

// func show(val any) {
// 	i, ok := val.(int)
// 	if ok {
// 		i = i + 2
// 		fmt.Println(i)
// 	}

// 	j, k := val.(string)
// 	if k {
// 		j = j + "hello"
// 		fmt.Println(j)
// 	}

// 	switch v := val.(type) {
// 	case int:
// 		i := v + 2
// 		fmt.Println(i)
// 	case string:
// 		j := v + "hello"
// 		fmt.Println(j)
// 	default:
// 		fmt.Println("not handle type.")
// 	}
// }

// type course struct {
// 	name, instructor string
// 	price            int
// }

// func (c *course) discount(d int) int {
// 	c.price = c.price - d
// 	fmt.Println("discount:", c.price)
// 	return c.price
// }

// type course struct {
// 	name, instructor string
// 	price            int
// }

// func (c *course) discount(d int) int {
// 	c.price = c.price - d
// 	fmt.Println("discount:", c.price)
// 	return c.price
// }

// func (c course) info() {
// 	fmt.Println("name:", c.name)
// 	fmt.Println("instructor:", c.instructor)
// 	fmt.Println("price:", c.price)
// }

// func (m movie) info() {
// 	fmt.Println("Output:")
// 	fmt.Printf("%s (%d) - %.2f\n", m.title, m.year, m.rating)
// 	fmt.Println("Genres:")
// 	for i := 0; i < len(m.genres); i++ {
// 		fmt.Printf("\t%v\n", m.genres[i])
// 	}
// }

// func adder() (func() int, func() int) {
// 	sum := 0
// 	return func() int {
// 			sum = sum + 1
// 			return sum
// 		}, func() int {
// 			return sum
// 		}
// }

// func show(sk []string) {
// 	l := len(sk)
// 	fmt.Printf("length: %d -- show: %#v\n", l, sk)
// }

// type movie struct {
// 	title       string
// 	year        int
// 	rating      float64
// 	genres      []string
// 	isSuperHero bool
// }
