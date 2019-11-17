package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shake, not stirred",
			"You",
			"In ssss",
		},
	}
	u2 := user{
		First: "Miss",
		Last:  "asca",
		Age:   32,
		Sayings: []string{
			"James",
			"scooff",
			"acddd",
		},
	}
	u3 := user{
		First: "M",
		Last:  "lllkkk",
		Age:   54,
		Sayings: []string{
			"scascgthy",
			"wqwooi",
			"tytt",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println("we scs ddvs")
	}
}
