package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrBadRequest = fmt.Errorf("invalid request")
	ErrNotFound   = fmt.Errorf("not found")
)

type serverHandler struct {
	db dbClient
}

func (s *serverHandler) handleRequest(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")

	err := reqValidation(segments)
	if err != nil {
		switch {
		case errors.Is(err, ErrBadRequest):
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("something went wrong..."))
		}
		return
	}

	var body []byte
	switch len(segments) {
	case 2:
		body, err = handleUserList(s.db)
	case 3:
		body, err = handleUser(s.db, segments)
	case 4:
		body, err = handleCity(s.db, segments)
	case 5:
		body, err = handleBuilding(s.db, segments)
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

func handleCity(db dbClient, segments []string) ([]byte, error) {
	cityID, err := strconv.Atoi(segments[3])
	if err != nil {
		return nil, fmt.Errorf("atoi %v got %w", err)
	}
	return json.Marshal(db.getCity(segments[2], cityID))
}

func handleBuilding(db dbClient, segments []string) ([]byte, error) {
	cityID, err := strconv.Atoi(segments[3])
	if err != nil {
		return nil, fmt.Errorf("atoi city ID %v got %w", err)
	}
	buildingID, err := strconv.Atoi(segments[3])
	if err != nil {
		return nil, fmt.Errorf("atoi building ID %v got %w", err)
	}
	b, err := db.getBuilding(segments[2], cityID, buildingID)
	if err != nil {
		return nil, err
	}
	return json.Marshal(b)
}

func reqValidation(segments []string) error {
	if len(segments) < 2 {
		return fmt.Errorf("%w: excpected at least 2 elements in path", ErrBadRequest)
	}

	for i, segment := range segments {
		switch i {
		case 1:
			if segment != "users" {
				return ErrBadRequest
			}
		case 2:
			// we expect a string, and it's a URL encoded string, so we're safe (ish) ?
		case 3, 4:
			// we expect numbers here
			_, err := strconv.Atoi(segment)
			if err != nil {
				return fmt.Errorf("%w: expecting a number here")
			}
		default:
			return fmt.Errorf("%w: too many elements in path", ErrBadRequest)
		}
	}
	return nil
}

func (s *serverHandler) init_db() mockDB {

	var database mockDB
	database.dbmock = s.db.getUserList()

	return database
}

func (s *serverHandler) handleRoot(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Welcome to Stickian!\n")
}
