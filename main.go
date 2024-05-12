// Reference: https://blog.logrocket.com/building-simple-app-go-postgresql/.
package main

import (
	"database/sql"
	"fmt"
	helper "go-dynamic-filters-perf/pkg"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	envFile, _ := godotenv.Read(".devcontainer/.env")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", envFile["POSTGRES_USER"], envFile["POSTGRES_PASSWORD"], envFile["POSTGRES_HOSTNAME"], envFile["POSTGRES_DB"])
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: uncomment if we want to re-setup.
	// helper.SetupTable(db)

	app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

// Routes.
func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	start := time.Now()

	queries := c.Queries()
	players := []helper.Player{}

	whereClauseArray := []string{}

	// Test fixed query.
	for k, v := range queries {
		whereClauseArray = append(whereClauseArray, fmt.Sprintf("%s='%s'", k, v))
	}

	// Test dynamic query.

	whereClause := ""
	// filterClause := ""
	if len(whereClauseArray) > 0 {
		whereClause = fmt.Sprintf("WHERE %s", strings.Join(whereClauseArray, " AND "))
		// filterClause = fmt.Sprintf("FILTER(%s) as referred FROM players", whereClause)
	}

	// query := fmt.Sprintf("SELECT COUNT(*) FROM players %s", filterClause)
	// fmt.Printf("%s\n", query)

	// totalRows, err := db.Query(query)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	c.JSON("An error occured")
	// }
	// defer totalRows.Close()

	// query = fmt.Sprintf("SELECT * FROM players LIMIT 10 %s", whereClause)
	// fmt.Printf("%s\n", query)

	query := fmt.Sprintf("SELECT * FROM players %s LIMIT 10", whereClause)
	fmt.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
		return c.JSON("An error occured")
	}
	defer rows.Close()

	for rows.Next() {
		var id uuid.UUID
		var clubId uuid.UUID
		var name string
		var country string
		var shirtName string

		rows.Scan(&id, &clubId, &name, &country, &shirtName)
		players = append(players, helper.Player{
			ID:        id,
			ClubId:    clubId,
			Name:      name,
			Country:   country,
			ShirtName: shirtName,
		})
	}

	end := time.Now()
	diff := end.UTC().Unix() - start.UTC().Unix()

	c.Append("response-time", fmt.Sprintf("%d", diff))

	return c.JSON(fiber.Map{
		"players": players,
	})
}
