package main
import (
    "fmt"
    "strings"
)


func main(){

	fmt.Println("Lets Do A Quiz!") 

	//NAME
	fmt.Printf("Enter your name: ") 
	var name string
	fmt.Scan(&name) //user input
    name = strings.Title(name)
	fmt.Printf("Your Name is %v \n", name)
	
	// AGE
	fmt.Printf("\nWhat is your age: ")
	var age int
	fmt.Scan(&age)

	//VERIFYING AGE
	if age > 10 {
		fmt.Println("Lets play")
	}else{
		fmt.Println("Too young, sorry")
		return //exit program if too young
	}


	//KEEPING SCORE
	score := 0

	//QUIZ STARTS
	fmt.Printf("Is Shawn Mendes an Actor or Singer? ")
	var ans string
	fmt.Scan(&ans)
	ans = strings.ToLower(ans)
	if ans == "singer"{
		fmt.Println("Correct")
		score++
	} else{
		fmt.Println("Wrong!!")
	}
	fmt.Printf("Is Orange a Fruit or Vegetable? ")
	
	fmt.Scan(&ans)
	ans = strings.ToLower(ans)
	if ans == "fruit"{
		fmt.Println("Correct")
		score++
	} else{
		fmt.Println("Wrong!!")
	}

	// CALCULATING
	percent := float64(score) / 2 * 100

	fmt.Printf("You scored %v out of 2.\nPercentage score: %.1f%%\n", score, percent)


}

//IN TERMINAL

//to run: go run quiz.go  
//to build executable file: go build quiz.go THEN ./quiz