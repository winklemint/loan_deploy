package main

import (
	cont "admin/Controller"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// func LeadtablePage(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "table-basic.html", r)
// 	//http.FileServer(http.Dir("form"))
// }

func main() {

	// config, err := con.LoadConfig("Config/config.yaml")
	// if err != nil {
	// 	log.Fatal("Failed to load config:", err)
	// }

	// db, err = con.ConnectDB(config)
	// if err != nil {
	// 	log.Fatal("Failed to connect to the database:", err)
	// }

	// defer db.Close()

	file, err := os.OpenFile("logrus.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logrus.log")
		panic(err)
	}
	logrus.SetOutput(file)
	logrus.SetLevel(logrus.TraceLevel)
	r := mux.NewRouter()

	protectedRoutes := r.PathPrefix("/").Subrouter()
	protectedRoutes.Use(cont.IsAuthenticated)

	// protectedRoutes2 := r.PathPrefix("/").Subrouter()
	// protectedRoutes2.Use(IsAuthenticated2)

	// Define routes
	protectedRoutes.PathPrefix("/form").Handler(http.StripPrefix("/form/", http.FileServer(http.Dir("form")))) //localhost:9000/form/ endpoint for accessing forms
	// r.HandleFunc("/index", TemplatePage).Methods("GET")

	r.HandleFunc("/proxy", cont.ProxyHandler).Methods("GET", "POST", "PATCH", "DELETE", "OPTIONS")

	// protectedRoutes.HandleFunc("/lead/table", LeadtablePage).Methods("GET")
	//r.HandleFunc("/landing", LandingPage).Methods("GET")
	r.HandleFunc("/login", cont.AdminLoginHandler).Methods("POST")
	r.HandleFunc("/subadmin/get", cont.GetAdmin).Methods("GET")
	r.HandleFunc("/get/admin", cont.GetAdminByID).Methods("GET")
	r.HandleFunc("/admin/logout", cont.LogOut).Methods("DELETE")
	r.HandleFunc("/admin/add", cont.AdminInsert).Methods("POST")
	r.HandleFunc("/admin/update/{id}", cont.UpadteAdmin).Methods("POST")
	r.HandleFunc("/admin/delete/{id}", cont.AdminDelete).Methods("DELETE")
	r.HandleFunc("/get/{id}", cont.AdminById).Methods("GET")

	// Sleep for a while to allow time for the API requests to complete
	//time.Sleep(5 * time.Second)
	// Create a reverse proxy for the React application
	// Register the /react route with the POST method and ProxyHandlerReact as the handler
	r.HandleFunc("/", cont.ProxyHandlerReact).Methods(http.MethodPost)

	// Create a reverse proxy for the React application
	reactURL, _ := url.Parse("http://localhost:3000")
	proxy := httputil.NewSingleHostReverseProxy(reactURL)

	// Register the reverse proxy as the default handler
	r.PathPrefix("/react").Handler(proxy)

	fmt.Println("Data fetched from the backend successfully.")

	http.ListenAndServe(":9000", r) // Start the server and listen for incoming requests
}
