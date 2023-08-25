package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/kthanasio/golang-server-api/configs"
	"github.com/kthanasio/golang-server-api/internal/entity"
	"github.com/kthanasio/golang-server-api/internal/infra/database"
	"github.com/kthanasio/golang-server-api/internal/pkg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	apiEndpoint string
	db          *gorm.DB
)

func main() {
	cfg, err := configs.LoadConfig("./")
	if err != nil {
		panic(err)
	}
	apiEndpoint = cfg.APIEndpoint

	// Open Database (5 seconds timeout)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Execute Database Migrations (5 seconds timeout)
	err = db.WithContext(ctx).AutoMigrate(&entity.Currency{})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/cotacao", HandleCotacao)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// HandleCotacao ...
func HandleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get Value from external API
	currency, err := pkg.GetTodayCurrency(ctx, apiEndpoint)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// Save result to Database
	currencyDB := database.NewCurrencyDB(db)
	err = currencyDB.Save(ctx, currency)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(strconv.FormatFloat(currency.Value, 'f', -1, 64)))
}
