// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	reservation "airbnb/features/reservation"

	mock "github.com/stretchr/testify/mock"
)

// ReservationServiceInterface is an autogenerated mock type for the ReservationServiceInterface type
type ReservationServiceInterface struct {
	mock.Mock
}

// CheckAvailability provides a mock function with given fields: input
func (_m *ReservationServiceInterface) CheckAvailability(input reservation.ReservationCore) (reservation.ReservationCore, error) {
	ret := _m.Called(input)

	var r0 reservation.ReservationCore
	var r1 error
	if rf, ok := ret.Get(0).(func(reservation.ReservationCore) (reservation.ReservationCore, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(reservation.ReservationCore) reservation.ReservationCore); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(reservation.ReservationCore)
	}

	if rf, ok := ret.Get(1).(func(reservation.ReservationCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewReservationServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewReservationServiceInterface creates a new instance of ReservationServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReservationServiceInterface(t mockConstructorTestingTNewReservationServiceInterface) *ReservationServiceInterface {
	mock := &ReservationServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
