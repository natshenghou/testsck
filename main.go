package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

type UserData struct{
	Id int
	CitizenId string
	Firstname string
	Lastname string
	BirthYear string
	FirstnameFather string
	FirstnameMother string
	LastnameFather string
	LastnameMother string
	SoldierId int
	AddressId int

}
func main(){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")

	if err != nil{
		fmt.Println("connect fail")
	}

	fmt.Println("connect sucsess")

	defer db.Close()

	results,_ :=db.Query("SELECT * FROM user")

	for results.Next() {
		var userData UserData
		err := results.Scan(
			&userData.Id, 
			&userData.CitizenId,
			&userData.Firstname, 
			&userData.Lastname,
			&userData.BirthYear, 
			&userData.FirstnameFather, 
			&userData.LastnameFather,
			&userData.FirstnameMother,
			&userData.LastnameMother,
			&userData.SoldierId,
			&userData.AddressId,
		) 
		if err != nil {
			panic(err.Error())
		}	
		fmt.Println(userData)
	}
}