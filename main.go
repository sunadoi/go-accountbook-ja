package main

import (
	"fmt"
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Item struct{
	category string
	price int
}

func main() {

	file, err := os.Create("accountbook.txt")
	// 開く場合にエラーが発生した場合
	if err != nil {
		// エラーを出力して終了する
		log.Fatal(err)
	}

	var n int
	fmt.Print("データの数を入力してください>")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if err := inputItem(file); err != nil {
			log.Fatal(err)
		}
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}

	if err := showItems(); err != nil {
		log.Fatal(err)
	}
}

func inputItem(file *os.File) error {
	var item Item

	fmt.Print("品目>")
	fmt.Scan(&item.category)

	fmt.Print("値段>")
	fmt.Scan(&item.price)

	line := fmt.Sprintf("%s %d\n", item.category, item.price)
	if _, err := file.WriteString(line); err != nil {
		return err
	}

	return nil
}

func showItems() error {
	file, err := os.Open("accountbook.txt")
	if err != nil {
		return err
	}

	fmt.Println("======================")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splited := strings.Split(line, " ")

		if len(splited) != 2 {
			return errors.New("パースに失敗しました")
		}

		category := splited[0]
		price, err := strconv.Atoi(splited[1])
		if err != nil {
			return err
		}

		fmt.Printf("%sに%d円使いました\n", category, price)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println("======================")

	return nil
}