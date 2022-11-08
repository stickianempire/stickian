package main

type user struct {
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
	getUserList() map[string][]user
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

func (m *mockDB) getUserList() map[string][]user {
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
}
