package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/rishikesh-revandikar/learning-g/04_Final_Basic_Project/internal/db"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Health check triggered", "method", r.Method)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	if db.DB == nil {
        slog.Error("Database connection is nil!")
        http.Error(w, "Database not initialized", http.StatusInternalServerError)
        return
    }
	switch r.Method {
	case http.MethodGet:
		rows, _ := db.DB.Query("SELECT id, name, grade FROM students")
		var students []db.Student
		for rows.Next() {
			var s db.Student
			rows.Scan(&s.ID, &s.Name, &s.Grade)
			students = append(students, s)
		}
		slog.Info("Fetched students", "count", len(students))
		json.NewEncoder(w).Encode(students)

	case http.MethodPost:
		var s db.Student
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res, _ := db.DB.Exec("INSERT INTO students (name, grade) VALUES (?, ?)", s.Name, s.Grade)
		id, _ := res.LastInsertId()
		s.ID = int(id)
		slog.Info("Student created", "id", s.ID, "name", s.Name)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(s)

	default:
		// Handling PUT and DELETE via path params
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 3 {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		id, _ := strconv.Atoi(pathParts[2])

		if r.Method == http.MethodPut {
			var s db.Student
			json.NewDecoder(r.Body).Decode(&s)
			db.DB.Exec("UPDATE students SET name=?, grade=? WHERE id=?", s.Name, s.Grade, id)
			slog.Info("Student updated", "id", id)
			w.WriteHeader(http.StatusNoContent)
		} else if r.Method == http.MethodDelete {
			db.DB.Exec("DELETE FROM students WHERE id=?", id)
			slog.Info("Student deleted", "id", id)
			w.WriteHeader(http.StatusNoContent)
		}
	}
}