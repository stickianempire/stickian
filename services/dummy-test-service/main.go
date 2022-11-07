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

	"github.com/gorilla/mux"
)

/*type user struct {
	UserID   int
	Username string
	Nickname string
	City     []city
}

type city struct {
	CityID    int
	Name      string
	Buildings []building
}

type building struct {
	BuildingID int
	Name       string
	Level      int
	Available  bool
}

type dbClient interface {
	getDB() map[string][]user
	getUser(username string) user
	getCity(username string, cityid int) city
	getBuilding(username string, cityid int, buildingid int) building
}
type mockDB struct {
	dbmock map[string][]user
}

type serverHandler struct {
	db dbClient
}

func (m *mockDB) getDB() map[string][]user {
	m.dbmock = map[string][]user{
		"Users": {
			{
				UserID:   69,
				Username: "MCKarol123",
				Nickname: "MCKarol",
				City: []city{
					{
						CityID: 1233,
						Name:   "Cidade da Karol",
						Buildings: []building{
							{
								BuildingID: 1,
								Name:       "City Hall",
								Level:      5,
								Available:  true,
							},
							{
								BuildingID: 2,
								Name:       "Barracks",
								Level:      2,
								Available:  true,
							},
						},
					},
				},
			},
			{
				UserID:   70,
				Username: "MCKevinho",
				Nickname: "KevinhoDiStreet",
				City: []city{
					{
						CityID: 1234,
						Name:   "Cidade do Kevinho",
						Buildings: []building{
							{
								BuildingID: 1,
								Name:       "City Hall",
								Level:      9,
								Available:  true,
							},
							{
								BuildingID: 2,
								Name:       "Barracks",
								Level:      5,
								Available:  true,
							},
							{
								BuildingID: 3,
								Name:       "Forge",
								Level:      0,
								Available:  true,
							},
							{
								BuildingID: 4,
								Name:       "Wall",
								Level:      0,
								Available:  false,
							},
						},
					},
				},
			},
		},
	}

	return m.dbmock
}

func (m *mockDB) getUser(username string) user {

	Data := m.dbmock
	var user_aux user

	for i := 0; i < len(Data["Users"]); i++ {
		//fmt.Printf("%v", Data["Users"][i])
		if username == Data["Users"][i].Username {
			user_aux = Data["Users"][i]
		}
	}

	return user_aux
}

func (m *mockDB) getCity(username string, cityid int) city {
	Data := m.dbmock
	var city_aux city

	for i := 0; i < len(Data["Users"]); i++ {
		for j := 0; j < len(Data["Users"][i].City); j++ {
			if (username == Data["Users"][i].Username) && (cityid == Data["Users"][i].City[j].CityID) {
				city_aux = Data["Users"][i].City[j]
			}
		}
	}
	return city_aux
}

func (m *mockDB) getBuilding(username string, cityid int, buildingid int) building {
	Data := m.dbmock
	var building_aux building

	for i := 0; i < len(Data["Users"]); i++ {
		for j := 0; j < len(Data["Users"][i].City); j++ {
			for k := 0; k < len(Data["Users"][i].City[j].Buildings); k++ {
				if (username == Data["Users"][i].Username) && (cityid == Data["Users"][i].City[j].CityID) &&
					(buildingid == Data["Users"][i].City[j].Buildings[k].BuildingID) {
					building_aux = Data["Users"][i].City[j].Buildings[k]
				}
			}
		}
	}
	return building_aux
}*/

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Welcome to Stickian!\n")
}

func (s *serverHandler) handleDB(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(s.db.getDB())
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	//return
}

func (s *serverHandler) handleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(s.db.getUser(vars["username"]))
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	//return
}

func (s *serverHandler) handleCity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	var_aux1, err := strconv.Atoi(vars["cityid"])
	if err != nil {
		fmt.Printf("Error converting String to Int\n")
		panic(err)
	}

	jsonResp, err := json.Marshal(s.db.getCity(vars["username"], var_aux1))
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	//return
}

func (s *serverHandler) handleBuilding(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	var_aux := vars["cityid"]

	var_aux1, err := strconv.Atoi(var_aux)
	if err != nil {
		fmt.Printf("Error converting String to Int\n")
		panic(err)
	}

	var_aux2 := vars["buildingid"]

	var_aux3, err := strconv.Atoi(var_aux2)
	if err != nil {
		fmt.Printf("Error converting String to Int\n")
		panic(err)
	}

	jsonResp, err := json.Marshal(s.db.getBuilding(vars["username"], var_aux1, var_aux3))
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	//return
}

func (s *serverHandler) init_db() mockDB {

	var database mockDB
	database.dbmock = s.db.getDB()

	return database
}

func main() {

	server := serverHandler{
		db: &mockDB{},
	}

	//user_test := "MCKarol123"
	//city_test := "1233"
	//building_test := "1"

	server.init_db()

	client, ctx, cancel, err1 := connect("mongodb://localhost:27017")
	if err1 != nil {
		panic(err1)
	}

	ping(client, ctx)

	r := mux.NewRouter()

	r.HandleFunc("/", getRoot)
	r.HandleFunc("/users/", server.handleDB)
	r.HandleFunc("/users/{username}/", server.handleUser)
	r.HandleFunc("/users/{username}/{cityid}/", server.handleCity)
	r.HandleFunc("/users/{username}/{cityid}/{buildingid}/", server.handleBuilding)
	http.Handle("/", r)

	err := http.ListenAndServe(":4000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}

	close(client, ctx, cancel)
}
