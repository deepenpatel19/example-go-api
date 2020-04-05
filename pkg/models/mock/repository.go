// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	context "context"
	models "github.com/cobbinma/example-go-api/pkg/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreatePet mocks base method
func (m *MockRepository) CreatePet(ctx context.Context, pet *models.Pet) models.PetError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePet", ctx, pet)
	ret0, _ := ret[0].(models.PetError)
	return ret0
}

// CreatePet indicates an expected call of CreatePet
func (mr *MockRepositoryMockRecorder) CreatePet(ctx, pet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePet", reflect.TypeOf((*MockRepository)(nil).CreatePet), ctx, pet)
}