package service

import (
	"context"
	"encoding/json"
)

func serveCompleteCityVew(ctx context.Context, db cityViewDatabaseClient, cityID int32) ([]byte, error) {
	cityInfo, err := db.GetCityInfo(ctx, cityID)
	if err != nil {
		return nil, err
	}
	cityInfo.LOCATION_X, cityInfo.LOCATION_Y = int16(cityInfo.LOCATION_ID>>16), int16(cityInfo.LOCATION_ID)
	cityInfoBytes, err := json.Marshal(cityInfo)
	if err != nil {
		return nil, err
	}
	return cityInfoBytes, nil
}

func serveBuildings(ctx context.Context, db cityViewDatabaseClient, cityID int32) ([]byte, error) {
	buildings, err := db.GetBuildings(ctx, cityID)
	if err != nil {
		return nil, err
	}
	buildingsBytes, err := json.Marshal(buildings)
	if err != nil {
		return nil, err
	}
	return buildingsBytes, nil
}

func serveResources(ctx context.Context, db cityViewDatabaseClient, cityID int32) ([]byte, error) {
	resources, err := db.GetResources(ctx, cityID)
	if err != nil {
		return nil, err
	}
	resourcesBytes, err := json.Marshal(resources)
	if err != nil {
		return nil, err
	}
	return resourcesBytes, nil
}

func serveUnits(ctx context.Context, db cityViewDatabaseClient, cityID int32) ([]byte, error) {
	units, err := db.GetUnits(ctx, cityID)
	if err != nil {
		return nil, err
	}
	unitsBytes, err := json.Marshal(units)
	if err != nil {
		return nil, err
	}
	return unitsBytes, nil
}
