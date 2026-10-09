package main

import (
	"encoding/json"
	. "fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type ar struct{}

func (ar15 ar) Error() string {
	return "Неккоректный возраст пользователя"
}
func check_age(age15 int) error {
	if age15 > 150 || age15 < 0 {
		return ar{}
	}
	return nil
}
func loadUsers(filename string) ([]User, error) {
	data, err := os.ReadFile(filename)
	var users []User
	if err != nil {
		return nil, err
	}
	err1 := json.Unmarshal(data, &users)
	if err1 != nil {
		return nil, err1
	}
	return users, nil
}

func saveUsers(user []User, filename string) error {
	data, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func addUsers(user []User, name string, age int) []User {
	user = append(user, User{Name: name, Age: age})
	return user
}

func findUser(user []User, name string) (User, bool) {
	for i := 0; i < len(user); i++ {
		if user[i].Name == name {
			return user[i], true
		}
	}
	return User{}, false
}
func update(user []User, name string, newAge int) ([]User, bool) {
	for i := 0; i < len(user); i++ {
		if user[i].Name == name {
			user[i].Age = newAge
			return user, true
		}
	}
	return user, false
}
func deleteUser(user []User, name string) ([]User, bool) {
	for i := 0; i < len(user); i++ {
		if user[i].Name == name {
			user = append(user[:i], user[i+1:]...)
			return user, true
		}
	}
	return user, false
}
func main() {
	var filename string
	Println("Введите имя файла с расширением json")
	Scan(&filename)
	user, err := loadUsers(filename)
	if err != nil {
		Println("Alarm !", err)
		return
	}

	for {
		Println("Выберите действие")
		Println("1.Показать всех пользователей")
		Println("2.Добавить пользователя")
		Println("3.Найти пользователя")
		Println("4.Изменить возраст")
		Println("5.Удалить пользователя")
		Println("0.Выход")
		var variant int
		Scan(&variant)
		switch variant {
		case 1:
			Println(user)
		case 2:
			var name15 string
			var age15 int
			Println("Введите имя")
			Scan(&name15)

			_, pravda := findUser(user, name15)
			if pravda == true {
				Println("Такой пользователь уже существует")
				continue
			}
			Println("Введите возраст")
			Scan(&age15)
			err := check_age(age15)
			if err != nil {
				Println(err)
				continue
			} else {
				user = addUsers(user, name15, age15)
				Println("Пользователь добавлен")
				err = saveUsers(user, filename)
				if err != nil {
					Println("alarm", err)
					return
				}
			}
		case 3:
			var name15 string
			Println("Введите имя")
			Scan(&name15)
			minUser, pravda := findUser(user, name15)
			if pravda != true {
				Println("Пользователь не найден")
			} else {
				Println(minUser)
			}
		case 4:
			var pravda bool
			var name15 string
			var age15 int
			Println("Введите имя")
			Scan(&name15)
			Println("Введите возраст")
			Scan(&age15)
			err := check_age(age15)
			if err != nil {
				Println(err)
				continue
			} else {
				user, pravda = update(user, name15, age15)
				if pravda != true {
					Println("Не найден")
				} else {
					Println("Возраст изменен")
				}
				err = saveUsers(user, filename)
				if err != nil {
					Println("alarm", err)
					return
				}
			}
		case 5:
			var name15 string
			var pravda bool
			Println("Введите имя")
			Scan(&name15)
			user, pravda = deleteUser(user, name15)
			if pravda != true {
				Println("Элемент не найден")
			} else {
				Println("Элемент удален")
			}
			err := saveUsers(user, filename)
			if err != nil {
				Println("alarm", err)
				return
			}
		case 0:
			Print("Возвращайтесь еще босс")
			return
		default:
			Println("Не туда ")
		}

	}
}
