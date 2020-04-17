package main

import "fmt"

type Item struct{
	category string
	price int
}

func main() {
	var n int
	fmt.Print("データの数を入力してください>")
	fmt.Scan(&n)

	items := make([]Item, 0, n)


	for i := 0; i < cap(items); i++ {
		items = inputItem(items)
	}

	showItems(items)
}

func inputItem(items []Item) []Item {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.category)

	fmt.Print("値段>")
	fmt.Scan(&item.price)

	items = append(items, item)

	return items
}

func showItems(items []Item) {
	fmt.Println("======================")
	for i := 0; i < len(items); i++ {
		fmt.Printf("%sに%d円使いました\n", items[i].category, items[i].price)
	}
	fmt.Println("======================")
}