// Code generated by mockery v2.0.3. DO NOT EDIT.

package mocks

import (
	eth "github.com/smartcontractkit/chainlink/core/services/eth"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"
)

// LogListener is an autogenerated mock type for the LogListener type
type LogListener struct {
	mock.Mock
}

// HandleLog provides a mock function with given fields: lb, err
func (_m *LogListener) HandleLog(lb eth.LogBroadcast, err error) {
	_m.Called(lb, err)
}

// JobID provides a mock function with given fields:
func (_m *LogListener) JobID() *models.ID {
	ret := _m.Called()

	var r0 *models.ID
	if rf, ok := ret.Get(0).(func() *models.ID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ID)
		}
	}

	return r0
}

// OnConnect provides a mock function with given fields:
func (_m *LogListener) OnConnect() {
	_m.Called()
}

// OnDisconnect provides a mock function with given fields:
func (_m *LogListener) OnDisconnect() {
	_m.Called()
}
