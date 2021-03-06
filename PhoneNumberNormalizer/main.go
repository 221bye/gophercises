package main

import (
	"fmt"
	"phnormal/data"
	phonedb "phnormal/db"
	"regexp"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = data.DbUser
	password = data.DbPassword
	dbname   = "gophercises_phone"
)

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	checkErr(phonedb.Reset("postgres", psqlInfo, dbname))

	checkErr(phonedb.Migrate("postgres", psqlInfo, dbname))

	db, err := phonedb.Open("postgres", psqlInfo)
	checkErr(err)
	defer db.Close()

	err = db.Populate()
	checkErr(err)

	phones, err := db.AllPhones()
	checkErr(err)

	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating/removing...", number)
			existing, err := db.FindPhone(number)
			checkErr(err)
			if existing != nil {
				checkErr(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				checkErr(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}

	phones, err = db.AllPhones()
	checkErr(err)

	fmt.Println("After normalizing:")
	for _, p := range phones {
		fmt.Printf("%+v\n", p)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}*/
