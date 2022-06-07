package service

import (
	"delos-farm-backend/domains"
	"delos-farm-backend/domains/mocks"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

var Farms = []domains.Farms{
	{ID: 1, Name: "Farm 1", Slug: "farm_1"},
	{ID: 2, Name: "Farm 2", Slug: "farm_2"},
}

var farmRepository = &mocks.FarmsRepositoryMock{Mock: mock.Mock{}}
var farmService = NewFarmsService(farmRepository)

func TestFarmsService_GetFound(t *testing.T) {
	farmRepository.Mock.On("Get", uint(1)).Return(Farms[0], nil)

	farm, err := farmService.Get(uint(1))
	assert.Nil(t, err, "should not return error")
	assert.NotNil(t, farm, "Farm should exist")

	assert.Equal(
		t,
		Farms[0].ID,
		farm.ID,
		fmt.Sprintf("Fetching wrong farm, id should be 1 but got %d", farm.ID),
	)
}

func TestFarmsService_GetNotFound(t *testing.T) {
	farmRepository.Mock.On("Get", uint(3)).Return(
		nil, errors.New("Farm not found"),
	)

	farm, err := farmService.Get(uint(3))
	assert.NotNil(t, err, "Should return error")

	isFarmEmpty := reflect.DeepEqual(domains.Farms{}, farm)
	assert.True(t, isFarmEmpty, "Farm object should be empty")
}

func TestFarmsService_CreateSuccess(t *testing.T) {
	newFarm := &domains.Farms{
		ID:   3,
		Name: "Farm 3",
		Slug: "farm_3",
	}
	farmRepository.Mock.On("Create", newFarm).Return(nil)

	err := farmService.Create(newFarm)
	assert.Nil(t, err, "should not return error")
}

func TestFarmsService_CreateDuplicate(t *testing.T) {
	farmRepository.Mock.On("Create", &Farms[0]).Return(
		errors.New("Farm already exists"),
	)

	err := farmService.Create(&Farms[0])
	assert.NotNil(t, err, "should return farm already exists error")
}

func TestFarmsService_DeleteSuccess(t *testing.T) {
	farmRepository.Mock.On("Delete", &Farms[0]).Return(nil)

	err := farmService.Delete(&Farms[0])
	assert.Nil(t, err, "should not return error")
}

func TestFarmsService_Delete(t *testing.T) {
	farmRepository.Mock.On("Delete", &Farms[0]).Return(nil)

	err := farmService.Delete(&Farms[0])
	assert.Nil(t, err, "should not return error")
}

func TestFarmsService_UpdateSuccess(t *testing.T) {
	updatedFarm := &domains.Farms{
		ID:   2,
		Name: "Farm 2 Updated",
		Slug: "farm_2_updated",
	}
	farmRepository.Mock.On("Update", updatedFarm).Return(nil)

	err := farmService.Update(updatedFarm)
	assert.Nil(t, err, "should not return error")
}

func TestFarmsService_UpdateAlreadyExists(t *testing.T) {
	farmRepository.Mock.On("Update", &Farms[0]).Return(
		errors.New("Farm already exists"),
	)

	err := farmService.Update(&Farms[0])
	assert.NotNil(t, err, "should return farm already exists error")
}

func TestFarmsService_GetAllSuccess(t *testing.T) {
	limit := "2"
	offset := "0"
	farmRepository.Mock.On("GetAll", 2, 0).Return(Farms, nil)

	farms, err := farmService.GetAll(limit, offset)
	assert.Nil(t, err, "should not return error")
	assert.NotNil(t, farms, "Farms should exist")
	assert.Equal(t, len(Farms), len(farms), "Must return 2 farms")
}

func TestFarmsService_GetAllNoData(t *testing.T) {
	limit := "2"
	offset := "5"
	farmRepository.Mock.On("GetAll", 2, 5).Return([]domains.Farms{}, nil)

	farm, err := farmService.GetAll(limit, offset)

	assert.NotNil(t, err, "Should not return no farms found error")
	assert.Equal(t, 0, len(farm), "Must contain 0 farms")
}
