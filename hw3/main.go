package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func main() {
	//task1()

	task2()
}

func task1() {

	for {
		fmt.Print("Ввод: ")
		var str string
		_, err := fmt.Scan(&str)
		if err != nil {
			fmt.Println(err)
			return
		}

		if str == "стоп" {
			fmt.Println("Выход")
			return
		}

		x, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Неверное значение")
			continue
		}

		fc := square(x)
		sc := multiplication(fc)
		fmt.Println("Произведение: ", <-sc)

	}
}

func square(x int) chan int {
	firstChan := make(chan int)
	go func() {
		res := x * x
		fmt.Println("Квадрат: ", res)
		firstChan <- res
	}()
	return firstChan
}

func multiplication(firstChan chan int) chan int {
	secondChan := make(chan int)
	x := <-firstChan
	go func() {
		res := x * 2
		secondChan <- res
	}()

	return secondChan
}

func task2() {
	var wg sync.WaitGroup
	wg.Add(1)
	go squaresNaturalNumbers(&wg)

	wg.Wait()
}

func squaresNaturalNumbers(wg *sync.WaitGroup) {
	i := 0
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		fmt.Println("Выхожу из программы")
		wg.Done()
	}()

	for {
		time.Sleep(time.Second)
		i++
		fmt.Println(i * i)
	}

}
