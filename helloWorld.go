package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	callSomeMethod()
	callSomeMethodWithMyName("TRINADH")
	declareVarNotWar()
	giveMeYourBioData()
}

func giveMeYourBioData() {
	var fName, lName string = "Trinadh", "Koya"
	salary := 8787
	var pinCodeArea int = 54545
	fmt.Println(fName + lName)
	fmt.Println("My First name is" + fName + "Last name is " + lName)
	fmt.Sprintf("My salary is %d", salary)
	fmt.Sprintf("PIN CODE IS %d", pinCodeArea)
}

func declareVarNotWar() {
	var firstName = "TRINADH KOYA"
	fmt.Println("Looks like you declared a variable and declared as well")
	checkWithTri(firstName)
}

func checkWithTri(name string) {
	fmt.Println("yes he declared and assigned his name with " + name)
}

func callSomeMethodWithMyName(s string) {
	fmt.Println("Yes it is me  " + s)
	fmt.Println("address of my name stored at", &s)
}

func callSomeMethod() {
	fmt.Println("I was called from main")
}
