// Code generated by MockGen. DO NOT EDIT.
// Source: ./accountRepo.go

// Package mock_account is a generated GoMock package.
package mock_account

import (
	domain "github.com/BrunoDM2943/pock_bank-api/domain"
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

// SaveAccount mocks base method
func (m *MockRepository) SaveAccount(account *domain.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAccount indicates an expected call of SaveAccount
func (mr *MockRepositoryMockRecorder) SaveAccount(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAccount", reflect.TypeOf((*MockRepository)(nil).SaveAccount), account)
}

// GetAccount mocks base method
func (m *MockRepository) GetAccount(ID int64) (*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ID)
	ret0, _ := ret[0].(*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockRepositoryMockRecorder) GetAccount(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockRepository)(nil).GetAccount), ID)
}
