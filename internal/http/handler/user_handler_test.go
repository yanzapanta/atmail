package handler

import (
	mock_service "atmail/internal/mock"
	"atmail/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/onsi/gomega"
)

func TestUserHandler_Get(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		httpStatus int
		err        error
	}{
		{name: "Get user successfully", id: "1", httpStatus: 200, err: nil},
		{name: "User not found", id: "100", httpStatus: 404, err: errors.New("user not found")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			ctrl := gomock.NewController(t)

			serviceMock := mock_service.NewMockUserService(ctrl)
			serviceMock.EXPECT().Get(gomock.Any()).Return(&model.User{
				ID:       1,
				Username: "username1",
				Email:    "email1",
				Age:      50,
			}, tt.httpStatus, tt.err).Times(1)

			handler := NewUserHandler(serviceMock)
			router := gin.New()
			router.GET("/users/:id", handler.Get)

			req, err := http.NewRequest(http.MethodGet, "/users/"+tt.id, nil)
			g.Expect(err).To(gomega.BeNil())
			writer := httptest.NewRecorder()
			router.ServeHTTP(writer, req)

			g.Expect(writer.Code).To(gomega.Equal(tt.httpStatus))
		})
	}
}

func TestUserHandler_Create(t *testing.T) {
	tests := []struct {
		name       string
		username   string
		email      string
		age        int
		httpStatus int
		err        error
	}{
		{name: "Create user successfully", httpStatus: 201, username: "username1", email: "email1@gmail.com", age: 34, err: nil},
		{name: "Email already exists", httpStatus: 400, username: "username1", email: "email1@gmail.com", age: 34, err: errors.New("email already exists")},
		{name: "Username already exists", httpStatus: 400, username: "username1", email: "email1@gmail.com", age: 34, err: errors.New("username already exists")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			ctrl := gomock.NewController(t)

			serviceMock := mock_service.NewMockUserService(ctrl)
			serviceMock.EXPECT().ValidateNewUser(gomock.Any()).Return(tt.err).Times(1)
			if tt.err == nil {
				serviceMock.EXPECT().Save(gomock.Any()).Return(&model.User{
					ID:       1,
					Username: tt.username,
					Email:    tt.email,
					Age:      tt.age,
				}, tt.err).Times(1)
			}

			handler := NewUserHandler(serviceMock)
			router := gin.New()
			router.POST("/users", handler.Create)
			var reqBytes []byte
			if tt.err == nil {
				r, err := json.Marshal(model.User{
					Username: tt.username,
					Email:    tt.email,
					Age:      tt.age,
				})
				g.Expect(err).To(gomega.BeNil())
				reqBytes = r
			}

			req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(reqBytes))
			g.Expect(err).To(gomega.BeNil())
			writer := httptest.NewRecorder()
			router.ServeHTTP(writer, req)

			g.Expect(writer.Code).To(gomega.Equal(tt.httpStatus))
		})
	}
}

func TestUserHandler_Update(t *testing.T) {
	tests := []struct {
		name       string
		id         uint
		username   string
		email      string
		age        int
		httpStatus int
		err        error
	}{
		{name: "Update user successfully", httpStatus: 200, id: 1, username: "username1", email: "email1@gmail.com", age: 34, err: nil},
		{name: "Email already exists", httpStatus: 400, id: 1, username: "username1", email: "email1@gmail.com", age: 34, err: errors.New("email already exists")},
		{name: "Invalid ID", httpStatus: 404, id: 100, username: "username1", email: "email1@gmail.com", age: 34, err: errors.New("no record found")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			ctrl := gomock.NewController(t)

			serviceMock := mock_service.NewMockUserService(ctrl)
			serviceMock.EXPECT().ValidateExistingUser(gomock.Any()).Return(tt.httpStatus, tt.err).Times(1)
			if tt.err == nil {
				serviceMock.EXPECT().Update(gomock.Any()).Return(&model.User{
					ID:       1,
					Username: tt.username,
					Email:    tt.email,
					Age:      tt.age,
				}, tt.err).Times(1)
			}

			handler := NewUserHandler(serviceMock)
			router := gin.New()
			router.PUT("/users/:id", handler.Update)
			var reqBytes []byte
			if tt.err == nil {
				r, err := json.Marshal(model.User{
					ID:       tt.id,
					Username: tt.username,
					Email:    tt.email,
					Age:      tt.age,
				})
				g.Expect(err).To(gomega.BeNil())
				reqBytes = r
			}
			idError := tt.err
			if idError != nil && idError.Error() == "no record found" {
				r, err := json.Marshal(model.User{
					Username: tt.username,
					Email:    tt.email,
					Age:      tt.age,
				})
				g.Expect(err).To(gomega.BeNil())
				reqBytes = r
			}

			req, err := http.NewRequest(http.MethodPut, "/users/"+strconv.Itoa(int(tt.id)), bytes.NewReader(reqBytes))
			g.Expect(err).To(gomega.BeNil())
			writer := httptest.NewRecorder()
			router.ServeHTTP(writer, req)

			g.Expect(writer.Code).To(gomega.Equal(tt.httpStatus))
		})
	}
}

func TestUserHandler_Delete(t *testing.T) {
	tests := []struct {
		name       string
		id         uint
		username   string
		email      string
		age        int
		httpStatus int
		err        error
	}{
		{name: "Delete user successfully", httpStatus: 200, id: 1, err: nil},
		{name: "Invalid ID", httpStatus: 404, id: 100, err: errors.New("no record found")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := gomega.NewWithT(t)
			ctrl := gomock.NewController(t)

			serviceMock := mock_service.NewMockUserService(ctrl)
			serviceMock.EXPECT().ValidateID(gomock.Any()).Return(tt.httpStatus, tt.err).Times(1)
			if tt.err == nil {
				serviceMock.EXPECT().Delete(gomock.Any()).Return(tt.err).Times(1)
			}

			handler := NewUserHandler(serviceMock)
			router := gin.New()
			router.DELETE("/users/:id", handler.Delete)

			req, err := http.NewRequest(http.MethodDelete, "/users/"+strconv.Itoa(int(tt.id)), nil)
			g.Expect(err).To(gomega.BeNil())
			writer := httptest.NewRecorder()
			router.ServeHTTP(writer, req)

			g.Expect(writer.Code).To(gomega.Equal(tt.httpStatus))
		})
	}
}
