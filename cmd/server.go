package cmd

import (
	"fmt"
	"main/db"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

var httpServerCommand = &cobra.Command{
	RunE:  runHTTPServer,
	Use:   "http_server",
	Short: "to run http server",
}

func runHTTPServer(_ *cobra.Command, _ []string) error {
	fmt.Println("==runHTTPServer()")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	gormDB, err := db.Connect()
	if err == nil {
		fmt.Println("Successfully connected to DB", gormDB)
	} else {
		panic("can't connect to DB")
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":3333", r)
	return nil
}
