// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "github.com/AJackTi/go-clean-architecture/internal/models"
	utils "github.com/AJackTi/go-clean-architecture/pkg/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockUseCase) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, user)
	ret0, _ := ret[0].(*models.UserWithToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockUseCaseMockRecorder) Register(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUseCase)(nil).Register), ctx, user)
}

// Login mocks base method
func (m *MockUseCase) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, user)
	ret0, _ := ret[0].(*models.UserWithToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockUseCaseMockRecorder) Login(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUseCase)(nil).Login), ctx, user)
}

// Update mocks base method
func (m *MockUseCase) Update(ctx context.Context, user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUseCaseMockRecorder) Update(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), ctx, user)
}

// Delete mocks base method
func (m *MockUseCase) Delete(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUseCaseMockRecorder) Delete(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), ctx, userID)
}

// GetByID mocks base method
func (m *MockUseCase) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, userID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockUseCaseMockRecorder) GetByID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUseCase)(nil).GetByID), ctx, userID)
}

// FindByName mocks base method
func (m *MockUseCase) FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", ctx, name, query)
	ret0, _ := ret[0].(*models.UsersList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName
func (mr *MockUseCaseMockRecorder) FindByName(ctx, name, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockUseCase)(nil).FindByName), ctx, name, query)
}

// GetUsers mocks base method
func (m *MockUseCase) GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*models.UsersList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, pq)
	ret0, _ := ret[0].(*models.UsersList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockUseCaseMockRecorder) GetUsers(ctx, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUseCase)(nil).GetUsers), ctx, pq)
}

// UploadAvatar mocks base method
func (m *MockUseCase) UploadAvatar(ctx context.Context, userID uuid.UUID, file models.UploadInput) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", ctx, userID, file)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadAvatar indicates an expected call of UploadAvatar
func (mr *MockUseCaseMockRecorder) UploadAvatar(ctx, userID, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockUseCase)(nil).UploadAvatar), ctx, userID, file)
}
