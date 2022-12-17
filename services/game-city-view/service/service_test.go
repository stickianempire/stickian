package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func unsafeMarshal(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

func Test_serveCompleteCityVew(t *testing.T) {
	randomErr := fmt.Errorf("random err")

	testcases := []struct {
		name           string
		mockCityInfo   *CityInfo
		mockError      error
		expectedResult []byte
		expectedError  error
	}{
		{
			name:           "happy path",
			mockCityInfo:   &CityInfo{},
			expectedResult: unsafeMarshal(&CityInfo{}),
		},
		{
			name:          "not found in db",
			mockError:     ErrNotFound,
			expectedError: ErrNotFound,
		},
		{
			name:          "random error from db is propagated",
			mockError:     randomErr,
			expectedError: randomErr,
		},
	}

	for _, test := range testcases {
		t.Run(test.name, func(tt *testing.T) {
			// given
			mockDB := mockCityViewDatabaseClient{
				getCityInfo: func(ctx context.Context, i int32) (*CityInfo, error) {
					return test.mockCityInfo, test.mockError
				},
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// when
			result, err := serveCompleteCityVew(ctx, &mockDB, 0)

			// then
			if !errors.Is(err, test.expectedError) {
				tt.Errorf("expected %v, got %v", test.expectedError, err)
			}
			if diff := cmp.Diff(result, test.expectedResult); diff != "" {
				tt.Errorf("unexpected result diff %v", diff)
			}
		})
	}
}
