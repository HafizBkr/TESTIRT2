package main

import (
	"fmt"
)

   func add(a int , b int){
	var c int
       c = a+b
	   fmt.Print("le resultat est ", c)
   }

func main (){
	fmt.Print("HELLO IRT2 \n")
	add(5,6)
}