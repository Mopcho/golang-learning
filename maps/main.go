package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int64) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		error := errors.New("Names has different length than phone numbers")
		return nil, error
	}
	userMap := make(map[string]user)
	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{phoneNumber: phoneNumbers[i], name: names[i]}
	}

	return userMap, nil
}

// don't touch below this line

type user struct {
	name        string
	phoneNumber int64
}

func test(names []string, phoneNumbers []int64) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int64{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int64{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int64{20955559812, 38385550982, 48265554567, 16045559873},
	)
}
