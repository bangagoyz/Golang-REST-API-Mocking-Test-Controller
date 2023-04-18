package service

import (
	"chapter3_2/models"
	"chapter3_2/repository/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestCarService_GetAll(t *testing.T) {
	carRepository := mocks.NewICarRepository(t)
	tests := []struct {
		name     string
		ss       *CarService
		want     []models.CarResponse
		mockFunc func()
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Case #1 - Success (Empty Data)",
			ss: &CarService{
				CarRepository: carRepository,
			},
			want: []models.CarResponse{},
			mockFunc: func() {
				carRepository.On("Get").Return([]models.Car{}, nil)
			},
			wantErr: false,
		},
		{
			name: "Case #2 - False",
			ss: &CarService{
				CarRepository: carRepository,
			},
			want: []models.CarResponse{
				{
					CarID:  "1",
					Title:  "Expander murah",
					Brand:  "mitsubishi",
					UserID: "1",
				},
				{
					CarID:  "2",
					Title:  "isuzu mu-x",
					Brand:  "isuzu",
					UserID: "1",
				},
			},
			mockFunc: func() {
				carRepository.On("Get").Return([]models.Car{
					{
						CarID:  "1",
						Title:  "Expander murah",
						Brand:  "mitsubishi",
						UserID: "1",
					},
					{
						CarID:  "2",
						Title:  "isuzu mu-x",
						Brand:  "isuzu",
						UserID: "2",
					},
				}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := tt.ss.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("CarService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarService.GetAll() = %v, want %v", got, tt.want)
			}
			tt.ss.CarRepository = carRepository
		})
	}
}

func TestCarService_GetOne(t *testing.T) {
	carRepository := mocks.NewICarRepository(t)
	type args struct {
		CarID string
	}
	tests := []struct {
		name     string
		ss       *CarService
		args     args
		want     models.CarResponse
		mockFunc func()
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Case #1 - Success",
			ss: &CarService{
				CarRepository: carRepository,
			},
			args: args{
				CarID: "1",
			},
			want: models.CarResponse{
				CarID:  "1",
				Title:  "isuzu mu-x",
				Brand:  "isuzu",
				UserID: "1",
			},
			mockFunc: func() {
				carRepository.On("GetOne", mock.Anything).Return(models.Car{
					CarID:  "1",
					Title:  "isuzu mu-x",
					Brand:  "isuzu",
					UserID: "1",
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "Case #2 - Not Found (Failed)",
			ss: &CarService{
				CarRepository: carRepository,
			},
			args: args{
				CarID: "1",
			},
			want: models.CarResponse{},
			mockFunc: func() {
				carRepository.On("GetOne", mock.Anything).Return(models.Car{}, models.ErrorNotFound)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := tt.ss.GetOne(tt.args.CarID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CarService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarService.GetOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
