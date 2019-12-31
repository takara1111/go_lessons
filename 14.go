package main 

import "fmt"

func main() {

	// score := 10

	if score := 70; score > 80 {
		fmt.Println("ぐれーと！")		
	} else if score > 60 {
		fmt.Println("ぐっど！")
	}	else {
		fmt.Println("うーん")
	}

}