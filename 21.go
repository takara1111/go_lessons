package main 

import "fmt"

func show (t interface{}) {
	// 型アサーション
	// _, ok := t.(japanese)
	// if ok {
	// 	fmt.Println("私は日本人")
	// } else {
	// 	fmt.Println("私は日本人じゃない")
	// }

	// 型スイッチ
	switch t.(type) {
	case japanese:
		fmt.Println("私は日本人")
	default:
		fmt.Println("私は日本人じゃない")
}
}

type greeter interface {
	greet()
}

type japanese struct {}
type american struct {}

func (j japanese) greet() {
	fmt.Println("こんちは")
}

func (a american) greet() {
	fmt.Println("はろー")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}