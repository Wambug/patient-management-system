package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/patienttracker/internal/services"
	"github.com/patienttracker/pkg/logger"
)

const version = "1.0.0"

type Server struct {
	Router   *mux.Router
	Services services.Service
	Log      *logger.Logger
}

func NewServer() *Server {
	mux := mux.NewRouter()
	conn := SetupDb("postgresql://postgres:secret@localhost:5432/patient_tracker?sslmode=disable")
	services := services.NewService(conn)
	logger := logger.New()
	server := Server{
		Router:   mux,
		Log:      logger,
		Services: services,
	}
	server.Routes()
	srve := http.Server{
		Addr:         "localhost:9000",
		Handler:      mux,
		ErrorLog:     log.New(logger, "", 0),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	fmt.Println("serving at port :9000")
	srve.ListenAndServe()
	return &server
}

func SetupDb(conn string) *sql.DB {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
	db.SetMaxOpenConns(65)
	db.SetMaxIdleConns(65)
	db.SetConnMaxLifetime(time.Hour)
	return db
}

func (server *Server) Routes() {
	http.Handle("/", server.Router)
	//server.Router.Use(server.contentTypeMiddleware)
	server.Router.HandleFunc("/v1/healthcheck", server.Healthcheck).Methods("GET")
	server.Router.HandleFunc("/v1/department", server.createdepartment).Methods("POST")
	server.Router.HandleFunc("/v1/department/{id:[0-9]+}", server.deletedepartment).Methods("DELETE")
	server.Router.HandleFunc("/v1/departments", server.findalldepartment).Methods("GET")
	server.Router.HandleFunc("/v1/department/{id:[0-9]+}", server.updatedepartment).Methods("POST")
	err := server.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
func (server *Server) Healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status: available\n")
	fmt.Fprintf(w, "version: %s\n", version)
	fmt.Fprintf(w, "Environment: Production")
}
