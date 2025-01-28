package main

import (
	"fmt"
)

   func add(a int , b int){
	var c int
       c = a+b
	   fmt.Print("le resultat est ", c)
   }


   func addd(b int , c int){
	var d int
       d = b+c
	   fmt.Print("le resultat est \n", d)
   }

func main (){
	fmt.Print("HELLO IRT2 \n")
	add(5,6)
	addd(19,2)

   func addd(a int , b int){
	var c int
       c = a+b
	   fmt.Print("le resultat est ", c)
   }


   func subs(c int , d int){
	var f int 
	 f= c-d
	 fmt.Print("\nle resultat de la soustraction est :", f)
   }
    
func main (){
	fmt.Print("HELLO IRT2 \n")
	add(5,6)
	addd(5,6)
	subs(6,5)

}