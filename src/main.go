package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandleFunc)
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("GET /public/", http.StripPrefix("/public/", fs))
	//
	LoadEnvFile()
	port := os.Getenv("OCI_PORT")
	intPort := os.Getenv("OCI_INT_PORT")
	if !IsRunningInDockerContainer() {
		intPort = port
	}
	srv := http.Server{
		Addr:    ":" + intPort,
		Handler: mux,
	}
	fmt.Println("Web Server is available at http://localhost:" + port)
	fmt.Println("Press Ctrl+C to stop")
	srv.ListenAndServe()
}

type HelloPageData struct {
	FullName  string
	CurrentDT string
}

func LoadEnvFile() string {
	// external configuration using environment variables
	var name = os.Getenv("ENV_NAME")
	if len(name) == 0 {
		name = ".env"
	}
	godotenv.Load(name)
	return name
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./resources/hello.html"))
	d := HelloPageData{FullName: "Everybody", CurrentDT: time.Now().Format(time.RFC3339)}
	err := tmpl.Execute(w, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func IsRunningInDockerContainer() bool {
	// docker creates a .dockerenv file at the root
	// of the directory tree inside the container.
	// if this file exists then the viewer is running
	// from inside a container so return true
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}
