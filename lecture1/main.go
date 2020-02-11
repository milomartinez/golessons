package main

import (
	"fmt"
	"log"
	"net/http"
	
)

func main () {

	/*
	//slide 2
	var variable typo 

	var x,y,z int
	var cadena string
	var boleano  bool*/

	//slide 3
	x:= 0
	cadena := "my string"
	booleano := true

	//no se puede hacer
	//cadena := "not my string" 
	//cadena = 0
	//cadena := 0

	//si se puede!!!
	cadena = "not my string" 



	fmt.Println("test 1",x,cadena,booleano);
	log.Println("test 2");


//slide 4
	istrue := false

	if istrue {
		fmt.Println("is true")
	} else {
		fmt.Println("is false")
	}


	//slide 5
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
		fmt.Println("sum:",sum);
	}


	myarray := []string{"duck", "cow", "dog"}
	for i,s  := range myarray {
		fmt.Println(i, s)
	}


	//slide 7
	http.HandleFunc("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "bar Hello, %s!", r.URL.Path[1:])
		fmt.Println(r);

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func fooHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	fmt.Println(r);
}