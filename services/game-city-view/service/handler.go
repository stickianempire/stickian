package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	idParam    = "id"
	scopeParam = "scope"

	completeScope = ""
	buildingScope = "buildings"
	resourceScope = "resources"
	unitsScope    = "units"

	requestTimeout = 5 * time.Second
)

var (
	ErrBadRequest = fmt.Errorf("invalid request")
	ErrNotFound   = fmt.Errorf("not found")
)

type ServerHandler struct {
	ctx context.Context
	db  cityViewDatabaseClient
}

func NewServerHandler(ctx context.Context, db cityViewDatabaseClient) *ServerHandler {
	return &ServerHandler{
		ctx: ctx,
		db:  db,
	}
}

// The city view handler will handle requests for city related information
// /city?id=0000							complete information
// /city?id=0000&scope=buildings			only building levels
// /city?id=0000&scope=resources			only resource ammounts
// /city?id=0000&scope=units				only unit amounts

func (s *ServerHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// create a context to this request handling
	ctx, cancel := context.WithTimeout(s.ctx, requestTimeout)
	defer cancel()

	resp, err := routeRequest(ctx, s.db, r.URL.RawQuery)

	switch {
	// Success case processing: 200 OK
	case err == nil:
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(resp)
		if err != nil {
			log.Printf("err writing %v", err)
		}
	// Requested something that did not exist: 404 NOT FOUND
	case errors.Is(err, ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
		_, err := fmt.Fprintf(w, "Not found\n")
		if err != nil {
			log.Printf("err writing %v", err)
		}
	// Miss-format on the request: 400 BAD REQUEST
	case errors.Is(err, ErrBadRequest):
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "Bad request\n")
		if err != nil {
			log.Printf("err writing %v", err)
		}
	// Unrecognized processing output: 500 INTERNAL SERVER ERROR
	default:
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprintf(w, "something went wrong...\n%v\n", err)
		if err != nil {
			log.Printf("err writing %v", err)
		}
	}
}

func routeRequest(ctx context.Context, db cityViewDatabaseClient, rawQuery string) ([]byte, error) {
	queryValues, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, fmt.Errorf("%w: on query parse got %v", ErrBadRequest, err)
	}
	cityIDstr := queryValues.Get(idParam)
	if cityIDstr == "" {
		return nil, fmt.Errorf("%w: city id necessary", ErrBadRequest)
	}

	cityID64, err := strconv.ParseInt(cityIDstr, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%w: city id %v got wrong format", ErrBadRequest, cityIDstr)
	}
	cityID := int32(cityID64)

	scope := queryValues.Get(scopeParam)
	switch scope {
	case completeScope:
		return serveCompleteCityVew(ctx, db, cityID)
	case buildingScope:
		return serveBuildings(ctx, db, cityID)
	case resourceScope:
		return serveResources(ctx, db, cityID)
	case unitsScope:
		return serveUnits(ctx, db, cityID)
	}
	return nil, ErrBadRequest
}
