package helper

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

func SetupTable(db *sql.DB) {
	toBeDroppedTables := []string{"players"}

	for _, table := range toBeDroppedTables {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
		if err != nil {
			panic(err)
		}
	}

	_, err := db.Exec(`
	CREATE TABLE players (
		id UUID PRIMARY KEY,
		club_id UUID,
	  name VARCHAR(255),
		country VARCHAR(255),
		shirt_name VARCHAR(255)
	)
	`)
	if err != nil {
		panic(err)
	}

	namesLength := len(names)
	countriesLength := len(countries)

	for it := 1; it < 11; it++ {
		fmt.Printf("Iteration %d\n", it)

		vals := []interface{}{}
		args := []string{}

		for i := 0; i < 10000; i++ {
			id := uuid.New().String()
			clubId := uuid.New().String()

			firstNameIdx := rand.Intn(namesLength)
			lastNameIdx := rand.Intn(namesLength)

			firstName := names[firstNameIdx][0]
			lastName := names[lastNameIdx][1]
			name := firstName + " " + lastName
			country := countries[rand.Intn(countriesLength)]
			shirtName := firstName

			if rand.Intn(2) == 1 {
				shirtName = lastName
			}

			argsIdx := i*5 + 1
			args = append(args, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", argsIdx, argsIdx+1, argsIdx+2, argsIdx+3, argsIdx+4))
			vals = append(vals, id, clubId, name, country, shirtName)
		}

		sqlStr := fmt.Sprintf("INSERT INTO players(id,club_id,name,country,shirt_name) VALUES %s", strings.Join(args, ","))

		stmt, _ := db.Prepare(sqlStr)

		_, err := stmt.Exec(vals...)
		if err != nil {
			panic(err)
		}
	}
}
