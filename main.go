package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, _ := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin) // Створюємо сканер для читання з консолі
	fmt.Println("Введіть текст:")
	text := "2"

	for text != "0" {
		file, _ := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		scanner.Scan() // Чекаємо введення користувача
		text = scanner.Text()
		if text == "1" {
			file, _ := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			os.Truncate("notes.txt", 0)
			defer file.Close()
		} else if text != "0" {
			file, _ := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			file.WriteString(text + "\n")
			defer file.Close()
		}
		data, _ := os.ReadFile("notes.txt")
		fmt.Println("\nВаші дані \n" + string(data))
		defer file.Close()

	}
	
	time.Sleep(3 * time.Second)

}

// )
// ЩОДЕННИк
// func main() {
// 	fmt.Println("Ви відкрили щоденник,оберіть дії : 1 - додати запис, 2 - переглянути, 3 - очисти все, 0 - вийти")
// 	var dia int
// 	var txt string
// 	fmt.Scan(&dia)
// 	file, err := os.OpenFile("diary.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
// 	if err != nil {
// 		fmt.Println("Помилка відкриття файлу:", err)
// 		return
// 	}
// 	defer file.Close()
// 	for dia != 0 {
// 		switch dia {
// 		case 1:
// 			file, _ := os.OpenFile("diary.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
// 			defer file.Close()
// 			fmt.Println("Введіть текст який хочете записати у щоденник : ")
// 			fmt.Scan(&txt)
// 			io.WriteString(file, txt+"\n")
// 			fmt.Println("Oберіть настпну дію")
// 			fmt.Scan(&dia)

// 		case 2:
// 			file, _ = os.Open("diary.txt")
// 			defer file.Close()
// 			fmt.Println("Вміст щоденника : ")
// 			content, _ := io.ReadAll(file)
// 			fmt.Println(string(content))
// 			fmt.Println("Oберіть настпну дію")
// 			fmt.Scan(&dia)

// 		case 3:
// 			file, _ = os.Open("diary.txt")
// 			os.Truncate("diary.txt", 0)
// 			fmt.Println("Файл очищено")
// 			fmt.Println("Oберіть настпну дію")
// 			fmt.Scan(&dia)
// 		}
// 	}
// 	fmt.Println("Дякую що скористались щоденником")
// 	time.Sleep(4 * time.Second

//
// s := []int{10, 20, 30, 40, 50}
// index := 2 // Видаляємо 30

// if index >= 0 && index < len(s) { // Перевірка, щоб не вийти за межі
// 	s = append(s[:index], s[index+1:]...)
// }

// fmt.Println(s) // Виведе: [10 20 40 50]

// Convertor
// func check(slice []string) (string, int) {
// 	var strc string
// 	for _, v := range slice {

// 		if v == "F" {
// 			num, _ := strconv.Atoi(strc)
// 			return v, num
// 		} else if v == "C" {
// 			num, _ := strconv.Atoi(strc)
// 			return "C", num

// 		}
// 		strc = strc + v

// 	}
// 	return " ", 0

// }

// func main() {
// 	var str string
// 	fmt.Println("Вітаємо у простому конвенторі градусів Цельсія у Фарінгейти та навпаки,введіть ваше значення за шаблоном(100C або 100F)")
// 	fmt.Scan(&str)
// 	slice := make([]string, len(str))
// 	for i, x := range str {
// 		slice[i] = string(x)

// 	}
// 	t, val := check(slice)
// 	if t == "F" {
// 		ct := "C"
// 		res := (val - 32) * 5 / 9
// 		fmt.Printf("Переконвертований результат : %d%s", res, ct)
// 	} else if t == "C" {
// 		ct := "F"
// 		res := val*9/5 + 32
// 		fmt.Printf("Переконвертований результат : %d%s", res, ct)
// 	} else {
// 		fmt.Println("Ви не ввели тип температури або ввели некоректно, перезавантажте програму")
// 	}

// 	time.Sleep(5 * time.Second)
// }

// anograma
// func check(slice []int, num int) bool {
// 	for _, v := range slice {
// 		if v == num {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	fmt.Println("Вітаємо у програмі яка преревіряє слова на анограмність,введіть перше слово :  ")
// 	var str1 string
// 	fmt.Scan(&str1)
// 	fmt.Println("Введіть друге слово :  ")
// 	var str2 string
// 	fmt.Scan(&str2)
// 	runes1 := []rune(str1)
// 	runes2 := []rune(str2)
// 	b := 0
// 	slice := make([]int, len(runes2))
// 	for n := range slice {
// 		slice[n] = -1
// 	}
// 	if len(runes1) == len(runes2) {
// 		for i := 0; i < len(runes1); i++ {
// 			if runes1[b] == runes2[i] && !check(slice, i) {
// 				slice[b] = i
// 				b++
// 				i = -1

// 			}
// 			if b == len(runes1) {
// 				fmt.Println("Ur words is anogram")
// 				break
// 			}
// 			if i == len(runes2)-1 {
// 				fmt.Println("Ur words is not anogram")
// 			}

// 		}
// 	} else {
// 		fmt.Println("Ur words is not anogram")
// 	}
// 	time.Sleep(5 * time.Second)

// }

// вгадування числа 7
// var c int64 = 0
// i := 0
// for c != 7 {
// 	i++
// 	c := rand.Intn(10)
// 	if c == 7 {
// 		break
// 	}um
// 	fmt.Printf("Спроба номер :  %v ,значення : %v\n", i, c)
// }
// fmt.Printf("Kількість спроб : %v\n", i)

// for s := 10; s >= 0; s-- {
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("\rДо виходу залишилось: %d секунд     ", s)
// }
