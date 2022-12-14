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
	getUser(username string) (user, error)
	getCity(username string, cityid int) (city, error)
	getBuilding(username string, cityid int, buildingid int) (building, error)
}

type mockDB struct {
	dbmock map[string][]user
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

func (m *mockDB) getUser(username string) (user, error) {
	Data := m.dbmock

	for i := 0; i < len(Data["Users"]); i++ {
		if username == Data["Users"][i].Username {
			return Data["Users"][i], nil
		}
	}

	return user{}, ErrNotFound
}

func (m *mockDB) getCity(username string, cityid int) (city, error) {
	Data := m.dbmock

	for i := 0; i < len(Data["Users"]); i++ {
		if Data["Users"][i].Username != username {
			continue
		}
		for j := 0; j < len(Data["Users"][i].City); j++ {
			if (username == Data["Users"][i].Username) && (cityid == Data["Users"][i].City[j].CityID) {
				return Data["Users"][i].City[j], nil
			}
		}
	}
	return city{}, ErrNotFound
}

func (m *mockDB) getBuilding(username string, cityid, buildingid int) (building, error) {
	Data := m.dbmock

	for i := 0; i < len(Data["Users"]); i++ {
		if Data["Users"][i].Username != username {
			continue
		}
		for j := 0; j < len(Data["Users"][i].City); j++ {
			if cityid != Data["Users"][i].City[j].CityID {
				continue
			}
			for k := 0; k < len(Data["Users"][i].City[j].Buildings); k++ {
				if buildingid == Data["Users"][i].City[j].Buildings[k].BuildingID {
					return Data["Users"][i].City[j].Buildings[k], nil
				}
			}
		}
	}
	return building{}, ErrNotFound
}
