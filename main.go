package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	var name string
	var result []User

	data, err := os.ReadFile("user.json")
	if err != nil {
		if os.IsNotExist(err) {
			result = []User{}
		} else {
			log.Fatalf("failed opening file: %s", err)
		}
	} else {
		err = json.Unmarshal(data, &result)
		if err != nil {
			log.Fatalf("failed unmarshalling file: %s", err)
		}
	}
	fmt.Print("Enter your name: ")
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()
	name = scaner.Text()

	fmt.Print("Сохранено!")

	fmt.Println("Все пользователи:")

	newID := 1
	if len(result) > 0 {
		newID = result[len(result)-1].ID + 1
	}

	users := User{
		ID:   newID,
		Name: name,
	}
	result = append(result, users)

	jsonData, err := json.MarshalIndent(result, "", "")
	os.WriteFile("user.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	for _, u := range result {
		fmt.Printf("%d - %s\n", u.ID, u.Name)
	}

}
