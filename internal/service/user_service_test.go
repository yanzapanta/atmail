package service

import (
	"atmail/internal/model"
	"atmail/internal/repository"
	"errors"
	"reflect"
	"testing"
)

type MockUser struct{}
type MockUserNotFound struct{}

func (u *MockUser) Get(id uint) (*model.User, error) {
	return &model.User{
		ID:       id,
		Username: "username1",
		Email:    "email1",
		Age:      60,
	}, nil
}

func (u *MockUserNotFound) Get(id uint) (*model.User, error) {
	return nil, errors.New("record not found")
}

func (u *MockUser) GetAll() (*[]model.User, error) {
	return &[]model.User{
		{
			ID:       1,
			Username: "username1",
			Email:    "email1",
			Age:      12,
		},
		{
			ID:       2,
			Username: "username2",
			Email:    "email2",
			Age:      34,
		},
	}, nil
}

func (u *MockUserNotFound) GetAll() (*[]model.User, error) {
	return nil, errors.New("no record found")
}

func (u *MockUser) GetUser(id uint) (*repository.User, error) {
	return &repository.User{
		ID:       id,
		Username: "username1",
		Email:    "email1",
		Age:      56,
	}, nil
}

func (u *MockUserNotFound) GetUser(id uint) (*repository.User, error) {
	return nil, errors.New("no record found")
}

func (u *MockUser) Delete(id uint) error {
	return nil
}

func (u *MockUserNotFound) Delete(id uint) error {
	return errors.New("user not found")
}

func (u *MockUser) IsEmailUnique(id *uint, email string) (bool, error) {
	return true, nil
}

func (u *MockUserNotFound) IsEmailUnique(id *uint, email string) (bool, error) {
	return false, errors.New("email already exists")
}

func (u *MockUser) IsUsernameUnique(id *uint, username string) (bool, error) {
	return true, nil
}

func (u *MockUserNotFound) IsUsernameUnique(id *uint, username string) (bool, error) {
	return false, errors.New("username already exists")
}

func (u *MockUser) Save(user repository.User) (*model.User, error) {
	return &model.User{
		ID:       1,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}, nil
}

func (u *MockUserNotFound) Save(user repository.User) (*model.User, error) {
	return &model.User{
		ID:       1,
		Username: "username1",
		Email:    "email1",
		Age:      10,
	}, nil
}

func (u *MockUser) Update(user repository.User) (*model.User, error) {
	return &model.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}, nil
}

func (u *MockUserNotFound) Update(user repository.User) (*model.User, error) {
	return nil, errors.New("user not found")
}

func Test_userService_Get(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		want1   int
		wantErr bool
	}{
		{
			name: "should return user successfully",
			fields: fields{
				userRepository: &MockUser{},
			},
			args: args{id: 1},
			want: &model.User{
				ID:       1,
				Username: "username1",
				Email:    "email1",
				Age:      60,
			},
			want1:   200,
			wantErr: false,
		},
		{
			name: "should fail to return user",
			fields: fields{
				userRepository: &MockUserNotFound{},
			},
			args:    args{id: 100},
			want:    nil,
			want1:   400,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepository: tt.fields.userRepository,
			}
			got, got1, err := u.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("userService.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_userService_GetAll(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]model.User
		wantErr bool
	}{
		{
			name: "should return users successfully",
			fields: fields{
				userRepository: &MockUser{},
			},
			want: &[]model.User{
				{
					ID:       1,
					Username: "username1",
					Email:    "email1",
					Age:      12,
				},
				{
					ID:       2,
					Username: "username2",
					Email:    "email2",
					Age:      34,
				},
			},
			wantErr: false,
		},
		{
			name: "should return users successfully",
			fields: fields{
				userRepository: &MockUserNotFound{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := u.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Save(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		req model.UserRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser *model.User
		wantErr  bool
	}{
		{
			name: "should create user successfully",
			fields: fields{
				userRepository: &MockUser{},
			},
			args: args{
				model.UserRequest{
					Username: "username1",
					Email:    "email1",
					Age:      10,
				},
			},
			wantUser: &model.User{
				ID:       1,
				Username: "username1",
				Email:    "email1",
				Age:      10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepository: tt.fields.userRepository,
			}
			gotUser, err := u.Save(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("userService.Save() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func Test_userService_Update(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		req model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "should update user successfully",
			fields: fields{
				userRepository: &MockUser{},
			},
			args: args{
				model.User{
					ID:       1,
					Username: "username10",
					Email:    "email10",
					Age:      15,
				},
			},
			want: &model.User{
				ID:       1,
				Username: "username10",
				Email:    "email10",
				Age:      15,
			},
			wantErr: false,
		},
		{
			name: "failed to update user",
			fields: fields{
				userRepository: &MockUserNotFound{},
			},
			args: args{
				model.User{
					ID:       1,
					Username: "username10",
					Email:    "email10",
					Age:      15,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := u.Update(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Delete(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should delete user successfully",
			fields: fields{
				userRepository: &MockUser{},
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "failed to delete user",
			fields: fields{
				userRepository: &MockUserNotFound{},
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepository: tt.fields.userRepository,
			}
			if err := u.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("userService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
