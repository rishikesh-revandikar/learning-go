package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/rishikesh-revandikar/learning-g/04_Final_Basic_Project/internal/db"
	"github.com/rishikesh-revandikar/learning-g/04_Final_Basic_Project/internal/routes"
)


func main(){
	logger := slog.New(slog.NewJSONHandler(os.Stdout,nil))
	slog.SetDefault(logger)

	db.InitDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", routes.HealthHandler)
	mux.HandleFunc("/students", routes.StudentsHandler)
	mux.HandleFunc("/students/", routes.StudentsHandler)

	wrappedMux := routes.LoggerMiddleware(mux)

	slog.Info("Server starting", "port", 8080)
	if err := http.ListenAndServe(":8080", wrappedMux); err != nil {
		slog.Error("Server failed", "error", err)
	}

}