package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var ErrNotFound = fmt.Errorf("not found")

type serverHandler struct {
	db dbClient
}

func (s *serverHandler) handleRequest(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")

	// url validate

	var body []byte
	var err error
	switch len(segments) {
	case 2:
		body, err = handleUserList(s.db)
	case 3:
		body, err = handleUser(s.db, segments)
	case 4:
		body, err = handleUserList(s.db)
	case 5:
		body, err = handleUserList(s.db)
	case 6:

	}

	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something went wrong..."))
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

	/* switch {
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
	} */
}

func handleUserList(db dbClient) ([]byte, error) {
	return json.Marshal(db.getUserList())
}

func handleUser(db dbClient, segments []string) ([]byte, error) {
	return json.Marshal(db.getUser(segments[2]))
}

func (s *serverHandler) handleCity(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	url := strings.Split(path, "/")

	mapdb := s.db.getUserList()

	flag := URLValidation(url, mapdb, w)

	if flag == 1 {
		url3, err1 := strconv.Atoi(url[3])
		if err1 != nil {
			log.Fatalf("Error happened in Atoi Err: %s", err1)
		}

		jsonResp, err := json.Marshal(s.db.getCity(url[2], url3))
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	} else {
		io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
	}
}

func (s *serverHandler) handleBuilding(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	url := strings.Split(path, "/")

	mapdb := s.db.getUserList()

	flag := URLValidation(url, mapdb, w)

	if flag == 1 {
		url3, err1 := strconv.Atoi(url[3])
		if err1 != nil {
			log.Fatalf("Error happened in Atoi Err: %s", err1)
		}

		url4, err2 := strconv.Atoi(url[4])
		if err2 != nil {
			log.Fatalf("Error happened in Atoi Err: %s", err2)
		}

		jsonResp, err := json.Marshal(s.db.getBuilding(url[2], url3, url4))
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	} else {
		io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
	}
}

func URLValidation(url []string, mapdb map[string][]user, w http.ResponseWriter) int {

	flag := 0

	users := url[1]

	switch {
	case len(url) == 4:
		{
			if users != "users" {
				io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
			} else {

				for i := 0; i < len(mapdb["Users"]); i++ {
					if url[2] == mapdb["Users"][i].Username {
						flag = 1
						break
					}
				}
			}
		}
	case len(url) == 5:
		{
			if users != "users" {
				io.WriteString(w, "Looks like you are trying to access an inexistent endpoint.\n")
				log.Fatalf("Error! Trying to access an inexistent endpoint.\n")
			} else {
				url3, err1 := strconv.Atoi(url[3])
				if err1 != nil {
					log.Fatalf("Error happened in Atoi Err: %s", err1)
				}

				for i := 0; i < len(mapdb["Users"]); i++ {
					if url[2] == mapdb["Users"][i].Username {
						for j := 0; j < len(mapdb["Users"][i].City); j++ {
							if url3 == mapdb["Users"][i].City[j].CityID {
								flag = 1
								break
							}
						}
					}
					if flag == 1 {
						break
					}
				}
			}
		}
	case len(url) == 6:
		{
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
							if flag == 1 {
								break
							}
						}
					}
					if flag == 1 {
						break
					}
				}
			}
		}
	default:
		break
	}
	return flag
}

func (s *serverHandler) init_db() mockDB {

	var database mockDB
	database.dbmock = s.db.getUserList()

	return database
}

func (s *serverHandler) handleRoot(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Welcome to Stickian!\n")
}
