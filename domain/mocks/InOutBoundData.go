// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "middleman-capstone/domain"

	mock "github.com/stretchr/testify/mock"
)

// InOutBoundData is an autogenerated mock type for the InOutBoundData type
type InOutBoundData struct {
	mock.Mock
}

// AddEntryData provides a mock function with given fields: newProduct
func (_m *InOutBoundData) AddEntryData(newProduct domain.InOutBounds) domain.InOutBounds {
	ret := _m.Called(newProduct)

	var r0 domain.InOutBounds
	if rf, ok := ret.Get(0).(func(domain.InOutBounds) domain.InOutBounds); ok {
		r0 = rf(newProduct)
	} else {
		r0 = ret.Get(0).(domain.InOutBounds)
	}

	return r0
}

// CekAdminEntry provides a mock function with given fields: newProduct
func (_m *InOutBoundData) CekAdminEntry(newProduct domain.InOutBounds) (bool, int, int) {
	ret := _m.Called(newProduct)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.InOutBounds) bool); ok {
		r0 = rf(newProduct)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(domain.InOutBounds) int); ok {
		r1 = rf(newProduct)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(domain.InOutBounds) int); ok {
		r2 = rf(newProduct)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// CekOwnerAdminEntry provides a mock function with given fields: newProduct
func (_m *InOutBoundData) CekOwnerAdminEntry(newProduct domain.InOutBounds) bool {
	ret := _m.Called(newProduct)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.InOutBounds) bool); ok {
		r0 = rf(newProduct)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CekOwnerEntry provides a mock function with given fields: newProduct
func (_m *InOutBoundData) CekOwnerEntry(newProduct domain.InOutBounds) bool {
	ret := _m.Called(newProduct)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.InOutBounds) bool); ok {
		r0 = rf(newProduct)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CekUserEntry provides a mock function with given fields: newProduct
func (_m *InOutBoundData) CekUserEntry(newProduct domain.InOutBounds) (bool, int, int) {
	ret := _m.Called(newProduct)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.InOutBounds) bool); ok {
		r0 = rf(newProduct)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(domain.InOutBounds) int); ok {
		r1 = rf(newProduct)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(domain.InOutBounds) int); ok {
		r2 = rf(newProduct)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// DeleteEntryAdminData provides a mock function with given fields: Productid
func (_m *InOutBoundData) DeleteEntryAdminData(Productid int) string {
	ret := _m.Called(Productid)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(Productid)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DeleteEntryUserData provides a mock function with given fields: productid, id
func (_m *InOutBoundData) DeleteEntryUserData(productid int, id int) string {
	ret := _m.Called(productid, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(productid, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ReadEntryAdminData provides a mock function with given fields: role
func (_m *InOutBoundData) ReadEntryAdminData(role string) []domain.InOutBounds {
	ret := _m.Called(role)

	var r0 []domain.InOutBounds
	if rf, ok := ret.Get(0).(func(string) []domain.InOutBounds); ok {
		r0 = rf(role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.InOutBounds)
		}
	}

	return r0
}

// ReadEntryUserData provides a mock function with given fields: id
func (_m *InOutBoundData) ReadEntryUserData(id int) []domain.InOutBounds {
	ret := _m.Called(id)

	var r0 []domain.InOutBounds
	if rf, ok := ret.Get(0).(func(int) []domain.InOutBounds); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.InOutBounds)
		}
	}

	return r0
}

// UpdateEntryAdminData provides a mock function with given fields: idproduct
func (_m *InOutBoundData) UpdateEntryAdminData(idproduct int) domain.InOutBounds {
	ret := _m.Called(idproduct)

	var r0 domain.InOutBounds
	if rf, ok := ret.Get(0).(func(int) domain.InOutBounds); ok {
		r0 = rf(idproduct)
	} else {
		r0 = ret.Get(0).(domain.InOutBounds)
	}

	return r0
}

// UpdateEntryUserData provides a mock function with given fields: idproduct, id
func (_m *InOutBoundData) UpdateEntryUserData(idproduct int, id int) domain.InOutBounds {
	ret := _m.Called(idproduct, id)

	var r0 domain.InOutBounds
	if rf, ok := ret.Get(0).(func(int, int) domain.InOutBounds); ok {
		r0 = rf(idproduct, id)
	} else {
		r0 = ret.Get(0).(domain.InOutBounds)
	}

	return r0
}

// UpdateQty provides a mock function with given fields: idcart, qty
func (_m *InOutBoundData) UpdateQty(idcart int, qty int) domain.InOutBounds {
	ret := _m.Called(idcart, qty)

	var r0 domain.InOutBounds
	if rf, ok := ret.Get(0).(func(int, int) domain.InOutBounds); ok {
		r0 = rf(idcart, qty)
	} else {
		r0 = ret.Get(0).(domain.InOutBounds)
	}

	return r0
}

// UpdateQtyAdminData provides a mock function with given fields: updatedData
func (_m *InOutBoundData) UpdateQtyAdminData(updatedData domain.InOutBounds) domain.InOutBounds {
	ret := _m.Called(updatedData)

	var r0 domain.InOutBounds
	if rf, ok := ret.Get(0).(func(domain.InOutBounds) domain.InOutBounds); ok {
		r0 = rf(updatedData)
	} else {
		r0 = ret.Get(0).(domain.InOutBounds)
	}

	return r0
}

// UpdateQtyUserData provides a mock function with given fields: updatedData
func (_m *InOutBoundData) UpdateQtyUserData(updatedData domain.InOutBounds) domain.InOutBounds {
	ret := _m.Called(updatedData)

	var r0 domain.InOutBounds
	if rf, ok := ret.Get(0).(func(domain.InOutBounds) domain.InOutBounds); ok {
		r0 = rf(updatedData)
	} else {
		r0 = ret.Get(0).(domain.InOutBounds)
	}

	return r0
}

type mockConstructorTestingTNewInOutBoundData interface {
	mock.TestingT
	Cleanup(func())
}

// NewInOutBoundData creates a new instance of InOutBoundData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInOutBoundData(t mockConstructorTestingTNewInOutBoundData) *InOutBoundData {
	mock := &InOutBoundData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}