package main

import "fmt"

func main() {
	connectToDB()
	defer disconnectFromDB()
	fmt.Printf("Do some transactions to Database\n")
}

func connectToDB() {
	fmt.Printf("Connection to Database!\n")
}

func disconnectFromDB() {
	fmt.Printf("Disconnection from Database!\n")
}
