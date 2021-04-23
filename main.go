package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// create employee struct based on data
// can be automatically paste the example json to get struct http://json2struct.mervine.net
type Employee struct {
	Domicile       domicile
	EmployeeNumber int64   `json:"employee_number"`
	Entity         string  `json:"entity"`
	Name           string  `json:"name"`
	Salary         float64 `json:"salary"`
}

type domicile struct {
	Country  string `json:"country"`
	IsRemote bool   `json:"is_remote"`
}

func main() {
	// TODO: task #1 - give me a skeleton!
	// fix: add comma after 1.5
	data := string(`
    {
        "name": "Golang",
        "entity": "Xendit",
        "employee_number": 10,
        "salary": 1.5,
        "domicile": {
            "country": "ID",
            "is_remote": true
        }
    }
    `)
	var employee Employee
	if err := json.Unmarshal([]byte(data), &employee); err != nil {
		fmt.Println("Task #1 failed!")
		return
	}

	// TODO: task #2 - I am a legal employee, include me into your database!
	// change from var database map[string]Employee into database := make(map[string]Employee)
	// because got an error panic: assignment to entry in nil map
	// required to initiliaze the map using make function (or map literal) before can add any elements
	database := make(map[string]Employee)
	database["Golang"] = Employee{
		EmployeeNumber: 10,
		Entity:         "Xendit",
		Name:           "Golang",
		Salary:         1.5,
		Domicile: domicile{
			Country:  "ID",
			IsRemote: true,
		},
	}

	if !reflect.DeepEqual(database["Golang"], employee) {
		fmt.Println("Task #2 failed!")
		return
	}
	return
}
