package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// to connect to the postgres databasemak
	_ "github.com/lib/pq"
)

type cityViewDatabaseClient interface {
	GetCityInfo(context.Context, int32) (*CityInfo, error)
	GetBuildings(context.Context, int32) (*Buildings, error)
	GetResources(context.Context, int32) (*Resources, error)
	GetUnits(context.Context, int32) (*Units, error)
}

type cityViewDatabaseClientImpl struct {
	db *sql.DB
}

func NewDatabaseClient(
	ctx context.Context,
	port int,
	host, user, password string,
) (*cityViewDatabaseClientImpl, error) {
	// TODO possibly use ssl mode and another dbname
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable", host, port, user, password)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &cityViewDatabaseClientImpl{db: db}, nil
}

func (c *cityViewDatabaseClientImpl) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

func (c *cityViewDatabaseClientImpl) GetCityInfo(ctx context.Context, cityID int32) (*CityInfo, error) {
	const query = `SELECT 
city_id,
player_id,
location_id,
city_name,
r_food_amount,
r_stick_amount,
r_rock_amount,
r_clay_amount,
r_coin_amount,
u_swordsman_amount,
u_stickoplite_amount,
u_pikeman_amount,
u_archer_amount,
u_crossbro_amount,
u_slinger_amount,
u_centaur_amount,
u_spidosaur_amount,
u_stickofants_amount,
u_trebuchet_amount,
u_ram_amount,
u_culverin_amount,
u_ciroustick_amount,
u_caravan_amount,
u_colonizer_amount,
u_fishboat_amount,
u_scientist_amount,
b_city_hall_level,
b_statue_level,
b_barracks_level,
b_wall_level,
b_quarry_level,
b_lumbermill_level,
b_clay_pit_level,
b_farm_level,
b_warehouse_level,
b_market_level,
b_fishing_spot_level,
b_bank_level,
b_embassy_level,
b_library_level,
b_research_lab_level,
b_hospital_level,
b_harbor_level,
b_airport_level,
b_mutant_lab_level,
b_temple_level,
income_tax FLOAT,
update_timestamp
FROM city
WHERE city_id=$1
`

	rows, err := c.db.QueryContext(ctx, query, cityID)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("closing rows %v", err)
		}
	}()

	cityInfoResult := make([]*CityInfo, 0, 1)

	for {
		if !rows.Next() {
			break
		}

		cityInfo := &CityInfo{}

		if err := rows.Scan(
			&cityInfo.CITY_ID,
			&cityInfo.PLAYER_ID,
			&cityInfo.LOCATION_ID,
			&cityInfo.CITY_NAME,
			&cityInfo.R_FOOD_AMOUNT,
			&cityInfo.R_STICK_AMOUNT,
			&cityInfo.R_ROCK_AMOUNT,
			&cityInfo.R_CLAY_AMOUNT,
			&cityInfo.R_COIN_AMOUNT,
			&cityInfo.U_SWORDSMAN_AMOUNT,
			&cityInfo.U_STICKOPLITE_AMOUNT,
			&cityInfo.U_PIKEMAN_AMOUNT,
			&cityInfo.U_ARCHER_AMOUNT,
			&cityInfo.U_CROSSBRO_AMOUNT,
			&cityInfo.U_SLINGER_AMOUNT,
			&cityInfo.U_CENTAUR_AMOUNT,
			&cityInfo.U_SPIDOSAUR_AMOUNT,
			&cityInfo.U_STICKOFANTS_AMOUNT,
			&cityInfo.U_TREBUCHET_AMOUNT,
			&cityInfo.U_RAM_AMOUNT,
			&cityInfo.U_CULVERIN_AMOUNT,
			&cityInfo.U_CIROUSTICK_AMOUNT,
			&cityInfo.U_CARAVAN_AMOUNT,
			&cityInfo.U_COLONIZER_AMOUNT,
			&cityInfo.U_FISHBOAT_AMOUNT,
			&cityInfo.U_SCIENTIST_AMOUNT,
			&cityInfo.B_CITY_HALL_LEVEL,
			&cityInfo.B_STATUE_LEVEL,
			&cityInfo.B_BARRACKS_LEVEL,
			&cityInfo.B_WALL_LEVEL,
			&cityInfo.B_QUARRY_LEVEL,
			&cityInfo.B_LUMBERMILL_LEVEL,
			&cityInfo.B_CLAY_PIT_LEVEL,
			&cityInfo.B_FARM_LEVEL,
			&cityInfo.B_WAREHOUSE_LEVEL,
			&cityInfo.B_MARKET_LEVEL,
			&cityInfo.B_FISHING_SPOT_LEVEL,
			&cityInfo.B_BANK_LEVEL,
			&cityInfo.B_EMBASSY_LEVEL,
			&cityInfo.B_LIBRARY_LEVEL,
			&cityInfo.B_RESEARCH_LAB_LEVEL,
			&cityInfo.B_HOSPITAL_LEVEL,
			&cityInfo.B_HARBOR_LEVEL,
			&cityInfo.B_AIRPORT_LEVEL,
			&cityInfo.B_MUTANT_LAB_LEVEL,
			&cityInfo.B_TEMPLE_LEVEL,
			&cityInfo.INCOME_TAX,
			&cityInfo.UPDATE_TIMESTAMP,
		); err != nil {
			return nil, err
		}
		cityInfoResult = append(cityInfoResult, cityInfo)
	}

	if len(cityInfoResult) > 1 {
		// should not happen
		return nil, fmt.Errorf("too many results")
	}

	if len(cityInfoResult) < 1 {
		return nil, ErrNotFound
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cityInfoResult[0], nil
}

func (c *cityViewDatabaseClientImpl) GetBuildings(ctx context.Context, cityID int32) (*Buildings, error) {
	const query = `SELECT 
b_city_hall_level,
b_statue_level,
b_barracks_level,
b_wall_level,
b_quarry_level,
b_lumbermill_level,
b_clay_pit_level,
b_farm_level,
b_warehouse_level,
b_market_level,
b_fishing_spot_level,
b_bank_level,
b_embassy_level,
b_library_level,
b_research_lab_level,
b_hospital_level,
b_harbor_level,
b_airport_level,
b_mutant_lab_level,
b_temple_level,
update_timestamp
FROM city
WHERE city_id=$1
`

	rows, err := c.db.QueryContext(ctx, query, cityID)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("closing rows %v", err)
		}
	}()

	buildingsResult := make([]*Buildings, 0, 1)

	for {
		if !rows.Next() {
			break
		}

		buildings := &Buildings{}

		if err := rows.Scan(
			&buildings.B_CITY_HALL_LEVEL,
			&buildings.B_STATUE_LEVEL,
			&buildings.B_BARRACKS_LEVEL,
			&buildings.B_WALL_LEVEL,
			&buildings.B_QUARRY_LEVEL,
			&buildings.B_LUMBERMILL_LEVEL,
			&buildings.B_CLAY_PIT_LEVEL,
			&buildings.B_FARM_LEVEL,
			&buildings.B_WAREHOUSE_LEVEL,
			&buildings.B_MARKET_LEVEL,
			&buildings.B_FISHING_SPOT_LEVEL,
			&buildings.B_BANK_LEVEL,
			&buildings.B_EMBASSY_LEVEL,
			&buildings.B_LIBRARY_LEVEL,
			&buildings.B_RESEARCH_LAB_LEVEL,
			&buildings.B_HOSPITAL_LEVEL,
			&buildings.B_HARBOR_LEVEL,
			&buildings.B_AIRPORT_LEVEL,
			&buildings.B_MUTANT_LAB_LEVEL,
			&buildings.B_TEMPLE_LEVEL,
			&buildings.UPDATE_TIMESTAMP,
		); err != nil {
			return nil, err
		}
		buildingsResult = append(buildingsResult, buildings)
	}

	if len(buildingsResult) > 1 {
		// should not happen
		return nil, fmt.Errorf("too many results")
	}

	if len(buildingsResult) < 1 {
		return nil, ErrNotFound
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return buildingsResult[0], nil
}

func (c *cityViewDatabaseClientImpl) GetResources(ctx context.Context, cityID int32) (*Resources, error) {
	const query = `SELECT 
r_food_amount,
r_stick_amount,
r_rock_amount,
r_clay_amount,
r_coin_amount,
update_timestamp
FROM city
WHERE city_id=$1
`

	rows, err := c.db.QueryContext(ctx, query, cityID)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("closing rows %v", err)
		}
	}()

	resourcesResult := make([]*Resources, 0, 1)

	for {
		if !rows.Next() {
			break
		}

		resources := &Resources{}

		if err := rows.Scan(
			&resources.R_FOOD_AMOUNT,
			&resources.R_STICK_AMOUNT,
			&resources.R_ROCK_AMOUNT,
			&resources.R_CLAY_AMOUNT,
			&resources.R_COIN_AMOUNT,
			&resources.UPDATE_TIMESTAMP,
		); err != nil {
			return nil, err
		}
		resourcesResult = append(resourcesResult, resources)
	}

	if len(resourcesResult) > 1 {
		// should not happen
		return nil, fmt.Errorf("too many results")
	}

	if len(resourcesResult) < 1 {
		return nil, ErrNotFound
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return resourcesResult[0], nil
}

func (c *cityViewDatabaseClientImpl) GetUnits(ctx context.Context, cityID int32) (*Units, error) {
	const query = `SELECT 
u_swordsman_amount,
u_stickoplite_amount,
u_pikeman_amount,
u_archer_amount,
u_crossbro_amount,
u_slinger_amount,
u_centaur_amount,
u_spidosaur_amount,
u_stickofants_amount,
u_trebuchet_amount,
u_ram_amount,
u_culverin_amount,
u_ciroustick_amount,
u_caravan_amount,
u_colonizer_amount,
u_fishboat_amount,
u_scientist_amount,
update_timestamp
FROM city
WHERE city_id=$1
`

	rows, err := c.db.QueryContext(ctx, query, cityID)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("closing rows %v", err)
		}
	}()

	unitsResult := make([]*Units, 0, 1)

	for {
		if !rows.Next() {
			break
		}

		units := &Units{}

		if err := rows.Scan(
			&units.U_SWORDSMAN_AMOUNT,
			&units.U_STICKOPLITE_AMOUNT,
			&units.U_PIKEMAN_AMOUNT,
			&units.U_ARCHER_AMOUNT,
			&units.U_CROSSBRO_AMOUNT,
			&units.U_SLINGER_AMOUNT,
			&units.U_CENTAUR_AMOUNT,
			&units.U_SPIDOSAUR_AMOUNT,
			&units.U_STICKOFANTS_AMOUNT,
			&units.U_TREBUCHET_AMOUNT,
			&units.U_RAM_AMOUNT,
			&units.U_CULVERIN_AMOUNT,
			&units.U_CIROUSTICK_AMOUNT,
			&units.U_CARAVAN_AMOUNT,
			&units.U_COLONIZER_AMOUNT,
			&units.U_FISHBOAT_AMOUNT,
			&units.U_SCIENTIST_AMOUNT,
			&units.UPDATE_TIMESTAMP,
		); err != nil {
			return nil, err
		}
		unitsResult = append(unitsResult, units)
	}

	if len(unitsResult) > 1 {
		// should not happen
		return nil, fmt.Errorf("too many results")
	}

	if len(unitsResult) < 1 {
		return nil, ErrNotFound
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return unitsResult[0], nil
}
