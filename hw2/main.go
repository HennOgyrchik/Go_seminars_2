package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	//task1()
	//task2()
	//task3()
	//task4_1()
	task4_2()
}

// -Напишите программу, которая на вход получала бы строку, введённую пользователем,
// а в файл писала № строки, дату и сообщение в формате:
// 2020-02-10 15:00:00 продам гараж.
// -При вводе слова exit программа завершает работу.
func task1() {
	file, err := os.Create("task1.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	i := 1
	for {
		fmt.Print("Введите строку: ")

		var str string
		_, err = fmt.Scan(&str)
		if err != nil {
			fmt.Print(err)
			return
		}

		if strings.EqualFold(str, "exit") {
			break
		}
		file.WriteString(fmt.Sprintf("%d) %s %s\n", i, time.Now().Format(time.DateTime), str))
		i++
	}

}

// Напишите программу, которая читает и выводит в консоль строки из файла,
// созданного в предыдущей практике, без использования ioutil.
// Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.
// Рекомендация:
// Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.
func task2() {
	file, err := os.Open("task1.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Print(err)
		return
	}

	buf := make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(string(buf))
}

// Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.
// Рекомендация:
// Для проверки создайте файл, установите режим только для чтения, закройте его, а затем,
// открыв, попытайтесь прочесть из него данные.
func task3() {
	file, err := os.Create("task3.txt")
	if err != nil {
		fmt.Print(err)
		return
	}

	if err := os.Chmod("task3.txt", 0444); err != nil {
		fmt.Print(err)
		file.Close()
		return
	}
	file.Close()

	file, err = os.Open("task3.txt")

	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Print(err)
		return
	}

	writer.WriteString("test")
	writer.Flush()

}

// Перепишите задачи 1 и 2, используя пакет ioutil.
func task4_1() {
	var b bytes.Buffer
	i := 1
	for {
		fmt.Print("Введите строку: ")

		var str string
		_, err := fmt.Scan(&str)
		if err != nil {
			fmt.Print(err)
			return
		}

		if strings.EqualFold(str, "exit") {
			break
		}
		b.WriteString(fmt.Sprintf("%d) %s %s\n", i, time.Now().Format(time.DateTime), str))
		i++
	}

	if err := ioutil.WriteFile("task4_1.txt", b.Bytes(), 0666); err != nil {
		fmt.Print(err)
		return
	}
}

func task4_2() {
	file, err := os.Open("task4_1.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()

	resBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(string(resBytes))

}
