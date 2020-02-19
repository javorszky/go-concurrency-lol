package main

import "fmt"

func f(i int) {
	p("Receiving", i)
}

func p(what string, i int) {
	fmt.Println(fmt.Sprintf("%s %d", what, i))
}

func main(){
	for i := 0; i < 40; i++ {
		p("Sending", i)
		go f(i)
	}
	fmt.Println("done")
}
