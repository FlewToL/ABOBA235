package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/FlewToL/myapp/utils"
	"github.com/FlewToL/myapp/internal/app/handlers"
)

type pingResp struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func withRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-Id")

		if id == "" {
			id = utils.NewID16()
		}

		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, r)
	})
}

func Run() {
	mux := http.NewServeMux()

	// Главная страница
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, Go project structure!")
	})

	mux.HandleFunc("/ping", handlers.Ping)

	mux.HandleFunc("/fail", func(w http.ResponseWriter, r *http.Request) {
		utils.LogRequest(r)

		utils.WriteErr(w, http.StatusBadRequest, "bad_request_example")
	})

	handler := withRequestID(mux)

	utils.LogInfo("Server is starting on :8080")

	if err := http.ListenAndServe(":8080", handler); err != nil {
		utils.LogError("server error: " + err.Error())
	}
}