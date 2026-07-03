package app

import (
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Controller struct {
	router      *http.ServeMux
	resourcesFs fs.FS
}

func NewController(router *http.ServeMux, resourcesFs fs.FS) (*Controller, error) {
	slog.Debug("ctrl::NewController() - Executing")
	ctrl := &Controller{router: router, resourcesFs: resourcesFs}
	//
	// add handlers
	router.Handle("GET /public/", http.FileServerFS(resourcesFs))
	router.HandleFunc("GET /", ctrl.defaultHandler)
	router.HandleFunc("GET /hello", ctrl.helloHandler)
	//
	return ctrl, nil
}

func (ctrl *Controller) defaultHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug("ctrl::defaultHandler() - Executing")
	http.Redirect(w, r, "/hello", http.StatusMovedPermanently)
}

func (ctrl *Controller) helloHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug("ctrl::helloHandler() - Executing")
	// processing input
	//
	// domain logic
	name := os.Getenv("FULLNAME")
	if len(name) == 0 {
		name = "Golang"
	}
	//
	// processing output (using templ)
	currentDT := time.Now().Format(time.RFC3339)
	err := HelloPage(name, currentDT).Render(r.Context(), w)
	if err != nil {
		slog.Error("ctrl::helloHandler() - Failed to process output", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
