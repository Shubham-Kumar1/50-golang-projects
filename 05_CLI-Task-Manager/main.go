// This is a CLI Task Manager Program in Go
// can be extended to multiple operations
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("****** Welcome to CLI Task Manager ******\n\n")
	var Tasks []string
	for {
		fmt.Println("Enter the Choice")
		fmt.Println("1.Add a Task\n2.Get all Tasks\n3.Exit")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("Enter the task to add : ")
			task := bufio.NewScanner(os.Stdin)
			task.Scan()
			Tasks = append(Tasks, task.Text())
		case 2:
			fmt.Println("***** Printing all the tasks *****\n")
			for i := 0; i < len(Tasks); i++ {
				fmt.Println(Tasks[i])
			}
			fmt.Println("**********************************")
		case 3:
			fmt.Println("*** Exiting form the Program ***")
			os.Exit(1)
		default:
			fmt.Println("Please Choose the correct option")
		}
	}
}
