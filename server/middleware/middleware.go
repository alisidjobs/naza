package middleware

import (
	"database/sql"
	"dnd/server/models"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	var err error
	// connection string
	connStr := "postgres://postgres:postgres@localhost?sslmode=disable"

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, notifying that we successfully connected to our database
	fmt.Println("You are now connected to the devdatabase database.")
}

func CreatePerso(w http.ResponseWriter, r *http.Request) {

}

func ModifyPerso(w http.ResponseWriter, r *http.Request) {

}

func DeletePerso(w http.ResponseWriter, r *http.Request) {

}

func DeleteAllPerso(w http.ResponseWriter, r *http.Request) {

}

func GetAllPerso(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Cool, your connected to the browser n")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM persos")

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	usrs := make([]models.Player, 0)

	for rows.Next() {
		usr := models.Player{}
		err := rows.Scan(&usr.ID, &usr.Name, &usr.Sexe)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		usrs = append(usrs, usr)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, usr := range usrs {
		fmt.Fprintf(w, "\n%d %s %B", usr.ID, usr.Name, usr.Sexe)
	}
}
