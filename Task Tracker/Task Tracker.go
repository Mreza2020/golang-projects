package main

import (
	"TaskTracker/Auxiliary_package"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func scan(m *string) {
	_, err := fmt.Scanln(m)
	Auxiliary_package.Error{Message: err}.ErrScan()

}
func scanText() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("err ReadString : %s ", err)
		return ""
	}
	text = strings.TrimSpace(text)
	return text
}

func main() {
	var (
		task            string
		name            string
		description     string
		status          string
		timeString      string
		output          string
		activity        string
		ok              string
		plans           string
		listPlans       []string
		allTask         []map[string]interface{}
		outputRemainder []Auxiliary_package.TaskAddSt
		outputDone      []Auxiliary_package.TaskAddSt
	)

	fmt.Println("Welcome to Task Tracker")
	output = Auxiliary_package.Open()
	if output == "ok" {
		fmt.Println("\nEnter your name?")
		name = scanText()
		for {
			fmt.Println("\nWhat is your plan today?")
			fmt.Println("\n * Add, Update, and Delete tasks == (Please enter the activity name) \n * Mark a task as in progress or done == 2 \n * List all tasks == 3 \n * List all tasks that are done == 4 \n * List all tasks that are not done == 5 ")
			scan(&task)

			switch task {
			case "add", "Add", "a", "A":
				fmt.Println("\nName of the activity?")
				activity = scanText()

				fmt.Println("\nBrief description of the activity?")
				description = scanText()
				fmt.Println("\nDid you do it?(yes , no)")
				scan(&status)
				if status == "yes" {
					fmt.Println("\nTime of activity?")
					timeString = scanText()
				} else {
					now := time.Now()
					timeString = now.Format("2006-01-02 15:04:05")
				}

				output = Auxiliary_package.TaskAdd(activity, description, status, timeString)
				fmt.Printf("%s Registration status? %s ", name, output)
			case "update", "Update", "u", "U":
				fmt.Println("\nName of the activity?")
				activity = scanText()
				fmt.Println("\nWas the activity successful?(yes , no)")
				scan(&status)

				if status == "yes" {
					output = Auxiliary_package.TaskUpdate(activity, status)
					fmt.Printf("%s Update status? %s ", name, output)
				} else {
					fmt.Println("\nWhenever this activity is completed, you can refer to it to accurately record daily activities.")
				}
			case "delete", "Delete", "d", "D":
				fmt.Println("\nName of the activity?")
				activity = scanText()
				fmt.Println("\nAre you sure you want to delete?")
				scan(&ok)
				if ok == "yes" {
					output = Auxiliary_package.TaskDelete(activity)
					fmt.Printf("%s Delete status? %s ", name, output)
				} else {
					fmt.Println("Cancel the erase operation")
				}
			case "2":
				fmt.Println("\nDo you want to add or check (y or n)?")
				scan(&ok)
				if ok == "y" {
					for {
						fmt.Println("\nWhich of your plans do you prioritize?")
						plans = scanText()
						//listPlans = strings.Split(plans, ",")
						listPlans = append(listPlans, plans)
						fmt.Println("\nContinue or stop (c or s)?")
						scan(&ok)
						if ok == "c" {
							continue
						} else {
							fmt.Println("___________________________________________________________")
							for id, list := range listPlans {
								fmt.Printf("\n \t* %d Plan Name: %s \n", id, list)
							}
							fmt.Println("___________________________________________________________")
							fmt.Println("\n!! Whenever you need the list: mark or Mark !!")
							break
						}
					}
				} else {
					fmt.Println(listPlans)
				}
			case "3":
				allTask = Auxiliary_package.AllTask()
				for id, task := range allTask {
					fmt.Printf("* %d", id)
					fmt.Printf("\t ID: %v\n", task["id"])
					fmt.Printf("\t Description: %v\n", task["description"])
					fmt.Printf("\t Status: %v\n", task["status"])
					fmt.Printf("\t Created At: %v\n", task["createdAt"])
					fmt.Printf("\t Updated At: %v\n", task["updatedAt"])
					fmt.Println("-----------")
				}
				fmt.Println("\n !! Finish !!")
			case "4":
				outputRemainder = Auxiliary_package.Remainder()
				fmt.Println("\nThe remaining activities are:")
				for id, taskR := range outputRemainder {
					fmt.Printf("* %d", id)
					fmt.Printf("\t ID: %v\n", taskR.Id)
					fmt.Printf("\t Description: %v\n", taskR.Description)
					fmt.Printf("\t Status: %v\n", taskR.Status)
					fmt.Printf("\t Created At: %v\n", taskR.CreatedAt)
					fmt.Printf("\t Updated At: %v\n", taskR.UpdatedAt)
					fmt.Println("-----------")
				}
			case "5":
				fmt.Println("\nAll activities carried out include:")
				outputDone = Auxiliary_package.Done()

				for id, taskD := range outputDone {
					fmt.Printf("* %d", id)
					fmt.Printf("\t ID: %v\n", taskD.Id)
					fmt.Printf("\t Description: %v\n", taskD.Description)
					fmt.Printf("\t Status: %v\n", taskD.Status)
					fmt.Printf("\t Created At: %v\n", taskD.CreatedAt)
					fmt.Printf("\t Updated At: %v\n", taskD.UpdatedAt)
					fmt.Println("-----------")
				}
			case "mark", "Mark":
				fmt.Println("___________________________________________________________")
				for id, list := range listPlans {
					fmt.Printf("\n \t* %d Plan Name: %s \n", id, list)
				}
				fmt.Println("___________________________________________________________")

			case "stop", "c":
				break

			default:
				fmt.Println("\nWe don't know that command")
			}
			time.Sleep(4 * time.Second)
		}

	} else {
		fmt.Println("\n!! Problem creating file !!")
		fmt.Println(output)
	}

}
