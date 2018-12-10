package main

import "fmt"

var g_grocery []string
var g_len int = 0 //total no of items in current array
func add_grocery(a ...string) {
	for _, d := range a {
		fmt.Println("Capacity:", cap(g_grocery))
		g_grocery = append(g_grocery, d)
	}

}
func list_grocery() {
	fmt.Println("Grocery List:")
	for i, d := range g_grocery {
		fmt.Println(i, d)
	}
}
func main() {
	add_grocery("Pear", "Banana", "Mango", "weed")
	/*add_grocery("banana")
	add_grocery("applet")
	add_grocery("apple")*/
	list_grocery()
}
