package main

import (
	"fmt"
	// "math/rand"
	// "time"
	// a2 "demo1/a2"
)

var p *int

func main() {
	a:=new(animal)
	a.age=33
	c:=new(cat)
	c.ani=a
	c.ani.age=1
	c.name="kat"
	fmt.Printf("%v", *c.ani)
	x:=c.eat()
	fmt.Printf("\n%v", x)
}


type animal struct {
	age int8
	length float32
 }
 type cat struct {
	name string
	ani *animal
 }
 func (c cat) eat() (x bool){
	fmt.Printf("%v eat fish",c)
	x=true
	return
 }