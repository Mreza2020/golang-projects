package main

import (
	"Number_Guessing_Game/Auxiliary_package"
	"Number_Guessing_Game/Auxiliary_package/Receiver"
	"fmt"
)

var (
	difficulty string
	dice       string
	dice1      int
	dice2      int
	n          int
	plural     int
)

func main() {
	fmt.Println("!!! Welcome to the game !!!")
	fmt.Println("You must add up the dice numbers to reach the game number. For example, if the game number is 20, the sum of your dice numbers must also be 20, neither more nor less.")
	fmt.Println("!!! Game difficulty !!!\nEasy\t1 \nNormal\t2 \nHard\t3 \nSuper Hard\t4 \nLegendary\t5 \nSuper Legendary\t6")
	Receiver.Receiver(&difficulty)
	switch difficulty {
	case "1":
		dice2 = Auxiliary_package.Rand(20)
		fmt.Printf("Goal number : %d\n", dice2)
		n += 22

		for i := 0; i < n; i++ {
			fmt.Println("Roll your dice (y)")
			Receiver.Receiver(&dice)
			fmt.Println("----------------------------------------------")
			if dice == "y" {
				dice1 = Auxiliary_package.Rand(6)
				plural += dice1
				dice = ""
				fmt.Printf("%-10s | %-10s | %-10s | %-10s\n", "Roll", "Similar", "Total", "Target")
				fmt.Printf("%-10d | %-10d | %-10d | %-10d\n", dice1, n-i, plural, dice2)
				if plural == dice2 {
					fmt.Printf("Congratulations, you won !!!\n")
					break
				}
				if plural > dice2 {
					plural -= plural
					fmt.Println("-------")
					fmt.Printf("Reset %d\n", plural)
					fmt.Println("-------")

				}

			} else {
				fmt.Println("Please type the word y.")
				fmt.Println("----------------------------------------------")

			}

		}
	case "2":
		dice2 = Auxiliary_package.Rand(42)
		fmt.Printf("Goal number : %d\n", dice2)
		n += 17

		for i := 0; i < n; i++ {
			fmt.Println("Roll your dice (y)")
			Receiver.Receiver(&dice)
			fmt.Println("----------------------------------------------")
			if dice == "y" {
				dice1 = Auxiliary_package.Rand(6)
				plural += dice1
				dice = ""
				fmt.Printf("%-10s | %-10s | %-10s | %-10s\n", "Roll", "Similar", "Total", "Target")
				fmt.Printf("%-10d | %-10d | %-10d | %-10d\n", dice1, n-i, plural, dice2)
				if plural == dice2 {
					fmt.Printf("Congratulations, you won !!!\n")
					break
				}
				if plural > dice2 {
					plural -= plural
					fmt.Println("-------")
					fmt.Printf("Reset %d\n", plural)
					fmt.Println("-------")

				}

			} else {
				fmt.Println("Please type the word y.")
				fmt.Println("----------------------------------------------")

			}

		}
	case "3":
		dice2 = Auxiliary_package.Rand(90)
		fmt.Printf("Goal number : %d\n", dice2)
		n += 25

		for i := 0; i < n; i++ {
			fmt.Println("Roll your dice (y)")
			Receiver.Receiver(&dice)
			fmt.Println("----------------------------------------------")
			if dice == "y" {
				dice1 = Auxiliary_package.Rand(6)
				plural += dice1
				dice = ""
				fmt.Printf("%-10s | %-10s | %-10s | %-10s\n", "Roll", "Similar", "Total", "Target")
				fmt.Printf("%-10d | %-10d | %-10d | %-10d\n", dice1, n-i, plural, dice2)
				if plural == dice2 {
					fmt.Printf("Congratulations, you won !!!\n")
					break
				}
				if plural > dice2 {
					plural -= plural
					fmt.Println("-------")
					fmt.Printf("Reset %d\n", plural)
					fmt.Println("-------")

				}

			} else {
				fmt.Println("Please type the word y.")
				fmt.Println("----------------------------------------------")

			}

		}
	case "4":
		dice2 = Auxiliary_package.Rand(200)
		fmt.Printf("Goal number : %d\n", dice2)
		n += 52

		for i := 0; i < n; i++ {
			fmt.Println("Roll your dice (y)")
			Receiver.Receiver(&dice)
			fmt.Println("----------------------------------------------")
			if dice == "y" {
				dice1 = Auxiliary_package.Rand(6)
				plural += dice1
				dice = ""
				fmt.Printf("%-10s | %-10s | %-10s | %-10s\n", "Roll", "Similar", "Total", "Target")
				fmt.Printf("%-10d | %-10d | %-10d | %-10d\n", dice1, n-i, plural, dice2)
				if plural == dice2 {
					fmt.Printf("Congratulations, you won !!!\n")
					break
				}
				if plural > dice2 {
					plural -= plural
					fmt.Println("-------")
					fmt.Printf("Reset %d\n", plural)
					fmt.Println("-------")

				}

			} else {
				fmt.Println("Please type the word y.")
				fmt.Println("----------------------------------------------")

			}

		}
	case "5":
		dice2 = Auxiliary_package.Rand(700)
		fmt.Printf("Goal number : %d\n", dice2)
		n += 160

		for i := 0; i < n; i++ {
			fmt.Println("Roll your dice (y)")
			Receiver.Receiver(&dice)
			fmt.Println("----------------------------------------------")
			if dice == "y" {
				dice1 = Auxiliary_package.Rand(6)
				plural += dice1
				dice = ""
				fmt.Printf("%-10s | %-10s | %-10s | %-10s\n", "Roll", "Similar", "Total", "Target")
				fmt.Printf("%-10d | %-10d | %-10d | %-10d\n", dice1, n-i, plural, dice2)
				if plural == dice2 {
					fmt.Printf("Congratulations, you won !!!\n")
					break
				}
				if plural > dice2 {
					plural -= plural
					fmt.Println("-------")
					fmt.Printf("Reset %d\n", plural)
					fmt.Println("-------")

				}

			} else {
				fmt.Println("Please type the word y.")
				fmt.Println("----------------------------------------------")

			}

		}
	case "6":
		dice2 = Auxiliary_package.Rand(1000)
		fmt.Printf("Goal number : %d\n", dice2)
		n += 250

		for i := 0; i < n; i++ {
			fmt.Println("Roll your dice (y)")
			Receiver.Receiver(&dice)
			fmt.Println("----------------------------------------------")
			if dice == "y" {
				dice1 = Auxiliary_package.Rand(6)
				plural += dice1
				dice = ""
				fmt.Printf("%-10s | %-10s | %-10s | %-10s\n", "Roll", "Similar", "Total", "Target")
				fmt.Printf("%-10d | %-10d | %-10d | %-10d\n", dice1, n-i, plural, dice2)
				if plural == dice2 {
					fmt.Printf("Congratulations, you won !!!\n")
					break
				}
				if plural > dice2 {
					plural -= plural
					fmt.Println("-------")
					fmt.Printf("Reset %d\n", plural)
					fmt.Println("-------")

				}

			} else {
				fmt.Println("Please type the word y.")
				fmt.Println("----------------------------------------------")

			}

		}
	}

}
