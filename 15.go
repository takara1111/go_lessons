package main 

import "fmt"

func main() {
	// signal := "red"
	// switch signal {
	// case "red":
	// 	fmt.Println("とまれ")
	// case "yellow":
	// 	fmt.Println("ちょっとまて")
	// case "blue", "green":
	// 	fmt.Println("すすめ")
	// default:
	// 	fmt.Println("おかしいぞ")
	// }

	score := 40
	switch {
	case score > 80:
		fmt.Println("さいこう")
	case score > 60:
		fmt.Println("いいぞ")
	default :
		fmt.Println("うーん")
	}
}