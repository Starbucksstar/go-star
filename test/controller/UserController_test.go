package controller_test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"star/src/controller"
	"star/src/entity"
	"testing"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) FindUser(user *entity.User) {
	args := m.Called(user)
	args.Get(0).(*entity.User).City = "wuhan"
	user.ID = 1
	user.City = "wuhan"
}

func (m *MockUserService) SaveUser(user entity.User) (uint, error) {
	args := m.Called(user)
	return 1, args.Error(1)
}

func (m *MockUserService) DeleteUser(user entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestUserController_Login(t *testing.T) {
	router := gin.Default()

	mockUserService := &MockUserService{}

	userController := controller.NewUserController(mockUserService)
	router.GET("/login", userController.Login)

	req, _ := http.NewRequest("GET", "/login?id=1&name=John&password=secret", nil)
	resp := httptest.NewRecorder()

	mockUserService.On("FindUser", mock.Anything).Return(&entity.User{Name: "John", Password: "123", City: "wuhan"}, nil)

	router.ServeHTTP(resp, req)

	var data map[string]interface{}
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	assert.Equal(t, "wuhan", data["user"].(map[string]interface{})["city"])
}
