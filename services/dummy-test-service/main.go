package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (s *serverHandler) getRoot(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	url := strings.Split(path, "/")

	switch {
	case len(url) <= 2:
		s.handleRoot(w, r)
	case len(url) == 3:
		s.handleUserList(w, r)
	case len(url) == 4:
		s.handleUser(w, r)
	case len(url) == 5:
		s.handleCity(w, r)
	case len(url) == 6:
		s.handleBuilding(w, r)
	default:
		io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
	}
}

func (s *serverHandler) handleUserList(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	url := strings.Split(path, "/")

	users := url[1]

	if users != "users" {
		io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
	} else {
		jsonResp, err := json.Marshal(s.db.getUserList())
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	}
}

func (s *serverHandler) handleUser(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	url := strings.Split(path, "/")

	users := url[1]

	mapdb := s.db.getUserList()

	if users != "users" {
		io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
	} else {
		flag := 0

		for i := 0; i < len(mapdb["Users"]); i++ {
			if url[2] == mapdb["Users"][i].Username {
				flag = 1
				break
			}
		}

		if flag == 1 {
			jsonResp, err := json.Marshal(s.db.getUser(url[2]))
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
		} else {
			io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
		}
	}
}

func (s *serverHandler) handleCity(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	url := strings.Split(path, "/")

	users := url[1]

	mapdb := s.db.getUserList()

	if users != "users" {
		io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
		log.Fatalf("Error! Trying to access an inexistent endpoint.\n")
	} else {
		url3, err1 := strconv.Atoi(url[3])
		if err1 != nil {
			log.Fatalf("Error happened in Atoi Err: %s", err1)
		}

		flag := 0

		for i := 0; i < len(mapdb["Users"]); i++ {
			if url[2] == mapdb["Users"][i].Username {
				for j := 0; j < len(mapdb["Users"][i].City); j++ {
					if url3 == mapdb["Users"][i].City[j].CityID {
						flag = 1
						break
					}
				}
			}
		}

		if flag == 1 {
			jsonResp, err := json.Marshal(s.db.getCity(url[2], url3))
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
		} else {
			io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
		}
	}
}

func (s *serverHandler) handleBuilding(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	url := strings.Split(path, "/")

	users := url[1]

	mapdb := s.db.getUserList()

	if users != "users" {
		io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
		log.Fatalf("Error! Trying to access an inexistent endpoint.\n")
	} else {
		url3, err1 := strconv.Atoi(url[3])
		if err1 != nil {
			log.Fatalf("Error happened in Atoi Err: %s", err1)
		}

		url4, err2 := strconv.Atoi(url[4])
		if err2 != nil {
			log.Fatalf("Error happened in Atoi Err: %s", err2)
		}

		flag := 0

		for i := 0; i < len(mapdb["Users"]); i++ {
			if url[2] == mapdb["Users"][i].Username {
				for j := 0; j < len(mapdb["Users"][i].City); j++ {
					if url3 == mapdb["Users"][i].City[j].CityID {
						for k := 0; k < len(mapdb["Users"][i].City[j].Buildings); k++ {
							if url4 == mapdb["Users"][i].City[j].Buildings[k].BuildingID {
								flag = 1
								break
							}
						}
					}
				}
			}
		}

		if flag == 1 {
			jsonResp, err := json.Marshal(s.db.getBuilding(url[2], url3, url4))
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
		} else {
			io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
		}
	}
}

func (s *serverHandler) init_db() mockDB {

	var database mockDB
	database.dbmock = s.db.getUserList()

	return database
}

func (s *serverHandler) handleRoot(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Welcome to Stickian!\n"))
}

func main() {

	server := serverHandler{
		db: &mockDB{},
	}

	server.init_db()

	client, ctx, cancel, err1 := connect("mongodb://localhost:27017")
	if err1 != nil {
		panic(err1)
	}

	ping(client, ctx)

	http.HandleFunc("/", server.getRoot)

	err := http.ListenAndServe(":4000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}

	close(client, ctx, cancel)
}
