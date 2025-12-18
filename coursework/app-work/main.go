package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}

var db *sql.DB

func main() {
	if err := initDB(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/events", eventsHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Сервер запущено на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() error {
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)
// postgresql://app:XgDBTEwQD6Hc74hZV6B3HN9zeN7vg1yQteCaWwBxrBsia4zYUoPQZTKli7YzYAfi@coursework-db-rw.coursework-ns.svc.cluster.local:5432/app		
// jdbc:postgresql://coursework-db-rw.coursework-ns:5432/app?password=XgDBTEwQD6Hc74hZV6B3HN9zeN7vg1yQteCaWwBxrBsia4zYUoPQZTKli7YzYAfi&user=app
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		return fmt.Errorf("Помилка читання БД: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("Помилка підключення до бази даних: %w", err)
	}
	fmt.Println("Успішно підключено до PostgreSQL через pgx/v5!")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT,
		start_time TIMESTAMP NOT NULL,
		end_time TIMESTAMP NOT NULL
	);`
	_, err = db.Exec(createTableQuery)
	return err
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		getEvents(w, r)
	case "POST":
		addEvent(w, r)
	default:
		http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
	}
}

func getEvents(w http.ResponseWriter, _ *http.Request) {
	rows, err := db.Query("SELECT id, title, description, start_time, end_time FROM events ORDER BY start_time ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.StartTime, &e.EndTime); err != nil {
			log.Println("Помилка сканування:", err)
			continue
		}
		e.Start, e.End = e.StartTime, e.EndTime
		events = append(events, e)
	}
	json.NewEncoder(w).Encode(events)
}

func addEvent(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO events(title, description, start_time, end_time) VALUES($1, $2, $3, $4) RETURNING id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var newID int
	err = stmt.QueryRow(
		event.Title,
		event.Description,
		event.StartTime,
		event.EndTime,
	).Scan(&newID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	event.ID = newID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
