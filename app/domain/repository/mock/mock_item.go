// Code generated by MockGen. DO NOT EDIT.
// Source: item.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/tuanna7593/gosample/app/domain/entity"
	repository "github.com/tuanna7593/gosample/app/domain/repository"
	valueobject "github.com/tuanna7593/gosample/app/domain/valueobject"
)

// MockItemRepository is a mock of ItemRepository interface.
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository.
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance.
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// AssignTx mocks base method.
func (m *MockItemRepository) AssignTx(txm repository.TransactionManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AssignTx", txm)
}

// AssignTx indicates an expected call of AssignTx.
func (mr *MockItemRepositoryMockRecorder) AssignTx(txm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignTx", reflect.TypeOf((*MockItemRepository)(nil).AssignTx), txm)
}

// Create mocks base method.
func (m *MockItemRepository) Create(ctx context.Context, item *entity.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockItemRepositoryMockRecorder) Create(ctx, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockItemRepository)(nil).Create), ctx, item)
}

// GetByID mocks base method.
func (m *MockItemRepository) GetByID(ctx context.Context, itemID valueobject.ItemID) (entity.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, itemID)
	ret0, _ := ret[0].(entity.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockItemRepositoryMockRecorder) GetByID(ctx, itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockItemRepository)(nil).GetByID), ctx, itemID)
}

// List mocks base method.
func (m *MockItemRepository) List(ctx context.Context, pagination valueobject.PaginationRequest) ([]entity.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, pagination)
	ret0, _ := ret[0].([]entity.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockItemRepositoryMockRecorder) List(ctx, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockItemRepository)(nil).List), ctx, pagination)
}

// Updates mocks base method.
func (m *MockItemRepository) Updates(ctx context.Context, item *entity.Item, values map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updates", ctx, item, values)
	ret0, _ := ret[0].(error)
	return ret0
}

// Updates indicates an expected call of Updates.
func (mr *MockItemRepositoryMockRecorder) Updates(ctx, item, values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updates", reflect.TypeOf((*MockItemRepository)(nil).Updates), ctx, item, values)
}
