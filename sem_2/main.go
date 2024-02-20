package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	//task1()
	//task2()
	//task3()
	//task4()
	//task5()
	//task6()
	task7()
}

func task1() {
	text := "Hello smth"
	file, err := os.Create("Hello.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла. ", err)
		return
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println(file.Name())
}

func task2() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла. ", err)
		return
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("Введите число от 1 до 100")
	file.WriteString("Введите число от 1 до 100\n")

	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			file.WriteString(fmt.Sprintf("Введено число %d", answer))
			if answer < 1 || answer > 100 {
				fmt.Println("Число должно быть в диапозоне от 1 до 100")
				file.WriteString("Число должно быть в диапозоне от 1 до 100\n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("Ура! Число угадано")
			file.WriteString("Ура! Число угадано")
			return
		} else {
			if answer < n {
				fmt.Println("Загаданное число больше")
				file.WriteString("Загаданное число больше\n")
			} else {
				fmt.Println("Загаданное число меньше")
				file.WriteString("Загаданное число меньше\n")
			}
		}
	}
}

func task3() {
	f, err := os.Open("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	//
	//buf := make([]byte, 256)
	//if _, err := io.ReadFull(f, buf); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Print("%s\n", buf)

	buf := make([]byte, 128)
	_, err = f.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func task4() {
	file, err := os.Create("some.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)

	defer file.Close()

	writer.WriteString("Say HI")
	writer.WriteString("\n")
	writer.WriteRune('a')
	writer.WriteString("\n")
	writer.WriteByte(67) //C
	writer.WriteString("\n")
	writer.Write([]byte{65, 66, 67}) //ABC
	writer.WriteString("\n")
	writer.Flush()
}

func task5() {
	fmt.Println("Введите имя пользователя")
	var username string
	fmt.Scan(&username)

	fmt.Println("Введите пароль")
	var pass string
	fmt.Scan(&pass)
	fmt.Println("Введите возраст")
	var age int
	fmt.Scan(&age)

	file, err := os.Create("task5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Ваш логин: %s\n", username))
	b.WriteString(fmt.Sprintf("Ваш пароль: %s\n", pass))
	b.WriteString(fmt.Sprintf("Ваш возраст: %d\n", age))
	_, err = file.Write(b.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func task6() {
	file, err := os.Create("some.txt")
	if err := os.Chmod("some.txt", 0444); err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer.WriteString("Say HI")
	writer.WriteString("\n")
	writer.WriteRune('a')
	writer.WriteString("\n")
	writer.WriteByte(67) //C
	writer.WriteString("\n")
	writer.Write([]byte{65, 66, 67}) //ABC
	writer.WriteString("\n")
	writer.Flush()
}

func task7() {
	var b bytes.Buffer

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("Введите число от 1 до 100")
	b.WriteString("Введите число от 1 до 100\n")

	for {
		var answer int
		for {
			_, _ = fmt.Scan(&answer)
			b.WriteString(fmt.Sprintf("Введено число %d\n", answer))
			if answer < 1 || answer > 100 {
				fmt.Println("Число должно быть в диапозоне от 1 до 100")
				b.WriteString("Число должно быть в диапозоне от 1 до 100\n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("Ура! Число угадано")
			b.WriteString("Ура! Число угадано")
			break
		} else {
			if answer < n {
				fmt.Println("Загаданное число больше")
				b.WriteString("Загаданное число больше\n")
			} else {
				fmt.Println("Загаданное число меньше")
				b.WriteString("Загаданное число меньше\n")
			}
		}
	}

	fileName := "log.txt"
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Лог: ")
	fmt.Println(string(resultBytes))
}
