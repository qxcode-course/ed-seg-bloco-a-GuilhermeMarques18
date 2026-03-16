package main

import "fmt"

func filter(age int) string {
	if age < 12 {
		return "crianca"
	} else if age < 18 {
		return "jovem"
	} else if age < 65 {
		return "adulto"
	} else if age < 1000 {
		return "idoso"
	} else {
		return "mumia"
	}

}
func main() {
	var name string
	fmt.Scanln(&name)
	var age int
	fmt.Scan(&age)

	fmt.Println(name, "eh", filter(age))
}
