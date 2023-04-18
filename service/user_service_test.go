package service

import (
	"chapter3_2/models"
	"reflect"
	"testing"
)

func TestUserService_Register(t *testing.T) {
	type args struct {
		userRegisterRequest models.UserRegisterRequest
	}
	tests := []struct {
		name    string
		us      *UserService
		args    args
		want    *models.UserRegisterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.us.Register(tt.args.userRegisterRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	type args struct {
		userLoginRequest models.UserLoginRequest
	}
	tests := []struct {
		name    string
		us      *UserService
		args    args
		want    models.UserLoginResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.us.Login(tt.args.userLoginRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
