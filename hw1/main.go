package main

import "fmt"

func main() {
	//Задание 1. Слияние отсортированных массивов
	//Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять
	////в один массив длиной девять.
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [5]int{5, 6, 7, 8, 9}

	fmt.Println(task1(arr1, arr2))

	//Задание 2. Сортировка пузырьком
	//Отсортируйте массив длиной шесть пузырьком.
	arr3 := [6]int{5, 1, 3, 6, 83, 19}
	fmt.Println(bubbleSort(arr3))
}

func task1(arr1 [4]int, arr2 [5]int) [9]int {
	var res [9]int

	for i, _ := range res {

		if i < 4 {
			res[i] = arr1[i]
		} else {
			res[i] = arr2[i-4]
		}
	}
	return res
}

func bubbleSort(arr [6]int) [6]int {
	size := len(arr)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
