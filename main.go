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
	fmt.Println(read(db))
	fmt.Println(readByCitizenId(db,"1600100386841"))
	fmt.Println(read(db))
	
}

func read(db *sql.DB) []UserData{

	results,_ :=db.Query("SELECT * FROM user")

	var userDataList []UserData

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
		userDataList = append(userDataList, userData)
	}
	return userDataList
}

func add(db *sql.DB) bool{
	statement,_ := db.Prepare(`INSERT INTO user 
		(citizen_id, firstname, lastname, birthyear, 
			firstname_father, lastname_father, firstname_mother,
			 lastname_mother, soldier_id,address_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`)
	
	defer statement.Close()
	_, err := statement.Exec("1552425252111","ชาติชาย2","เพ็ชรเม็ด","1985","สุชาติ","เพ็ชรเม็ด","แม่","เพ็ชรเม็ด","1","1")
		if err != nil {
			panic(err.Error())
			return false
		}
		return true
}

func remove(db *sql.DB,id string) bool{
	statement,_ := db.Prepare("DELETE FROM user WHERE user_id=?") 

	_,err := statement.Exec(id)
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}
func edit(db *sql.DB,id string,fartherName string) bool{
	statement,_ := db.Prepare("UPDATE `user` SET firstname_father=? WHERE user_id=?") 
	defer statement.Close()
	_, err := statement.Exec(fartherName,id)
		if err != nil {
			panic(err.Error())
			return false
		}
		return true
}

func readByCitizenId(db *sql.DB,citizenId string) UserData{

	results,_ :=db.Query("SELECT * FROM user WHERE citizen_id= ?",citizenId)

	var userData UserData

	for results.Next() {
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
	}
	return userData
}