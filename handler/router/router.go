package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// HealthzHandlerのインスタンスを作成してエンドポイントを登録
	health := handler.NewHealthzHandler()
	mux.HandleFunc("/healthz", health.ServeHTTP)

	return mux
}
