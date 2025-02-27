package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"prj0/domain"
	"sort"
	"strconv"
	"time"
)

const (
	totalPoints       = 100
	pointsPerQuestion = 20
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у грі!")

	users := getUsers()
	for _, user := range users {
		if user.Id >= id {
			id = user.Id + 1
		}
	}

	sortAndSave(users)
	for {
		menu()

		choise := ""
		fmt.Scan(&choise)

		switch choise {
		case "1":
			user := play()
			users = getUsers()
			users = append(users, user)
			sortAndSave(users)
		case "2":
			users = getUsers()
			for _, u := range users {
				fmt.Printf(
					"Id: %v, Name: %s, Time: %v\n",
					u.Id, u.Name, u.TimeSpent,
				)
			}
		case "4":
			return
		case "3":
			fmt.Println("Ви впевнені що хочете очистити рейтинг? 1 - так,2 - ні")
			rem := ""
			fmt.Scan(&rem)
			if rem == "1" {
				os.Truncate("users.json", 0)
				fmt.Println("Рейтинг очищено,оберіть наступну дію")
			} else if rem == "2" {
				fmt.Println("Дякую що одумались,оберіть наступну дію")

			} else {
				fmt.Println("Ви обрали не ту дію,спробуйте ще раз")

			}

		default:
		}
	}

}

func menu() {
	fmt.Println("1. Грати")
	fmt.Println("2. Рейтинг")
	fmt.Println("3. Очистити")
	fmt.Println("4. Вийти")
}

func play() domain.User {
	TimeStart := time.Now()
	myPoints := 0

	for myPoints < totalPoints {
		x, y := rand.Intn(2), rand.Intn(2)
		char := rand.Intn(4)
		chs := ""
		switch char {
		case 0:
			chs = "+"
			fmt.Printf("\n%v %s %v = ", x, chs, y)

			ans := ""
			fmt.Scan(&ans)
			if ans == "HESOYAM" {
				fmt.Println("А ви читер пане,але гаразд")
				myPoints = totalPoints
			}

			ansInt, err := strconv.Atoi(ans)

			if err != nil && ans != "HESOYAM" {
				fmt.Println("Невдале значення, давай по новой")
			} else if ans != "HESOYAM" {
				if ansInt == x+y {
					myPoints += pointsPerQuestion
					fmt.Printf("Правильно! У вас %v очок!", myPoints)
				} else {
					fmt.Println("Не праивльно!")
				}
			}
		case 1:
			chs = "-"
			fmt.Printf("\n%v %s %v = ", x, chs, y)

			ans := ""
			fmt.Scan(&ans)
			if ans == "HESOYAM" {
				fmt.Println("А ви читер пане,але гаразд")
				myPoints = totalPoints
			}

			ansInt, err := strconv.Atoi(ans)

			if err != nil && ans != "HESOYAM" {
				fmt.Println("Невдале значення, давай по новой")
			} else {
				if ansInt == x-y {
					myPoints += pointsPerQuestion
					fmt.Printf("Правильно! У вас %v очок!", myPoints)
				} else {
					fmt.Println("Не праивльно!")
				}
			}
		case 2:
			chs = "*"
			fmt.Printf("\n%v %s %v = ", x, chs, y)

			ans := ""
			fmt.Scan(&ans)
			if ans == "HESOYAM" {
				fmt.Println("А ви читер пане,але гаразд")
				myPoints = totalPoints
			}

			ansInt, err := strconv.Atoi(ans)

			if err != nil && ans != "HESOYAM" {
				fmt.Println("Невдале значення, давай по новой")
			} else {
				if ansInt == x*y {
					myPoints += pointsPerQuestion
					fmt.Printf("Правильно! У вас %v очок!", myPoints)
				} else {
					fmt.Println("Не праивльно!")
				}
			}
		case 3:
			chs = "/"
			if y == 0 {
				y = 1
			}
			fmt.Printf("\n%v %s %v = ", x, chs, y)

			ans := ""
			fmt.Scan(&ans)
			if ans == "HESOYAM" {
				fmt.Println("А ви читер пане,але гаразд")
				myPoints = totalPoints

			}

			ansInt, err := strconv.Atoi(ans)

			if err != nil && ans != "HESOYAM" {
				fmt.Println("Невдале значення, давай по новой")
			} else {
				if ansInt == x/y && y != 0 {
					myPoints += pointsPerQuestion
					fmt.Printf("Правильно! У вас %v очок!", myPoints)
				} else {
					fmt.Println("Не праивльно!,або вам не пощастило натрапити на Y = 0")
				}
			}

		}

	}

	TimeFinish := time.Now()
	timeSpent := TimeFinish.Sub(TimeStart)

	fmt.Printf("\nВаш час: %v", timeSpent)
	fmt.Print("Введіть ваше ім'я: ")

	name := ""
	fmt.Scan(&name)

	user := domain.User{
		Id:        id,
		Name:      name,
		TimeSpent: timeSpent,
	}
	id++

	return user
}

func sortAndSave(users []domain.User) {
	sort.SliceStable(users, func(i int, j int) bool {
		return users[i].TimeSpent < users[j].TimeSpent
	})

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Printf("sortAndSave(os.OpenFile): %s", err)
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("sortAndSave(file.Close()): %s", err)
		}
	}()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(users)
	if err != nil {
		log.Printf("sortAndSave(encoder.Encode)): %s", err)
		return
	}
}

func getUsers() []domain.User {
	var users []domain.User
	file, err := os.Open("users.json")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_, err = os.Create("users.jsom")
			if err != nil {
				log.Printf("getUsers(os.Create): %s", err)
			}
			return nil
		}
		log.Printf("getUsers(os.Open): %s", err)
		return nil
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		log.Printf("getUsers(decoder.Decode): %s", err)
	}

	return users
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
