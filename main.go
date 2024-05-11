// Reference: https://blog.logrocket.com/building-simple-app-go-postgresql/.
package main

import (
	"database/sql"
	"fmt"
	helper "go-dynamic-filters-perf/pkg"
	"log"
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Player struct {
	ID        string `json:"id"`
	ClubId    string `json:"clubId"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	ShirtName string `json:"shirtName"`
}

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	envFile, _ := godotenv.Read(".devcontainer/.env")
	dbName := envFile["POSTGRES_DB"]

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", envFile["POSTGRES_USER"], envFile["POSTGRES_PASSWORD"], envFile["POSTGRES_HOSTNAME"], dbName)
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	toBeDroppedTables := []string{"teams", "players"}

	for _, table := range toBeDroppedTables {

		_, err = db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
		if err != nil {
			panic(err)
		}
	}

	_, err = db.Exec(`
	CREATE TABLE players (
		id UUID PRIMARY KEY,
		clubId UUID,
	  name TEXT,
		country TEXT,
		shirtName TEXT
	)
	`)
	if err != nil {
		panic(err)
	}

	namesLength := len(helper.Names)
	countriesLength := len(helper.Countries)

	for i := 1; i < 100000; i++ {
		id := uuid.New()
		clubId := uuid.New()

		firstNameIdx := rand.Intn(namesLength)
		lastNameIdx := rand.Intn(namesLength)

		firstName := helper.Names[firstNameIdx][0]
		lastName := helper.Names[lastNameIdx][1]
		name := firstName + " " + lastName
		country := rand.Intn(countriesLength)
		shirtName := firstName

		if rand.Intn(2) == 1 {
			shirtName = lastName
		}

		db.Exec("INSERT into players VALUES ($1, $2, $3, $4, $5)", id, clubId, name, country, shirtName)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

// Routes.
func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	players := []string{}

	rows, err := db.Query("SELECT * FROM players")
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&res)
		players = append(players, res)
	}
	return c.JSON(fiber.Map{
		"players": players,
	})
}

type Todo struct {
	Action string `json:"action"`
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := Todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Action != "" {
		_, err := db.Exec("INSERT into players VALUES ($1)", newTodo.Action)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}

	return c.Redirect("/")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	olditem := c.Query("olditem")
	newitem := c.Query("newitem")
	db.Exec("UPDATE players SET item=$1 WHERE item=$2", newitem, olditem)
	return c.Redirect("/")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from players WHERE item=$1", todoToDelete)
	return c.SendString("deleted")
}
