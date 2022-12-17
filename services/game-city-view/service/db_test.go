package service

import "context"

// define the mock for testing this db in other places
type mockCityViewDatabaseClient struct {
	getCityInfo  func(context.Context, int32) (*CityInfo, error)
	getBuildings func(context.Context, int32) (*Buildings, error)
	getResources func(context.Context, int32) (*Resources, error)
	getUnits     func(context.Context, int32) (*Units, error)
}

func (c *mockCityViewDatabaseClient) GetCityInfo(ctx context.Context, id int32) (*CityInfo, error) {
	return c.getCityInfo(ctx, id)
}
func (c *mockCityViewDatabaseClient) GetBuildings(ctx context.Context, id int32) (*Buildings, error) {
	return c.getBuildings(ctx, id)
}
func (c *mockCityViewDatabaseClient) GetResources(ctx context.Context, id int32) (*Resources, error) {
	return c.getResources(ctx, id)
}
func (c *mockCityViewDatabaseClient) GetUnits(ctx context.Context, id int32) (*Units, error) {
	return c.getUnits(ctx, id)
}
