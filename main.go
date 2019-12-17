package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func swapStrings(s1, s2 string) (string, string) {
	return s2, s1
}

func isAFibonacciNumber(num int) bool {
	if num == 0 {
		return true
	}

	buf1, buf2, buf3 := 1, 1, 1
	for buf2 < num {
		buf3 = buf1 + buf2
		buf1 = buf2
		buf2 = buf3
	}

	return buf2 == num
}

func getCountListTo(num int) []int {
	l := make([]int, num, num)
	for i := 1; i <= num; i++ {
		l[i-1] = i
	}
	return l
}

func Pic(dx, dy int) (pic [][]uint8) {
	pic = make([][]uint8, dx, dx)
	for x := 0; x < dx; x++ {
		pic[x] = make([]uint8, dy, dy)
		for y := 0; y < dy; y++ {
			pic[x][y] = uint8(x + y)
		}
	}
	return
}

func WordCount(s string) (count map[string]int) {
	count = make(map[string]int)

	splittedString := strings.Fields(s)
	for _, word := range splittedString {
		_, hasWord := count[word]
		if hasWord {
			count[word]++
		} else {
			count[word] = 1
		}
	}
	return
}

func fibonacci() func() int {
	var seq = []int{0, 1}
	var i = 0

	var incIndex = func() {
		i++
	}

	return func() int {
		if i > 1 {
			nextElem := seq[i-2] + seq[i-1]
			seq = append(seq, nextElem)
		}

		defer incIndex()

		return seq[i]
	}
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	//fmt.Println("My favorite number is", rand.Intn(10))
	//fmt.Println(swapStrings("Hello", "World"))
	//pic.Show(Pic)
	//f := fibonacci()
	//
	//for i := 0; i< 10; i++ {
	//	fmt.Println(f())
	//}

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
