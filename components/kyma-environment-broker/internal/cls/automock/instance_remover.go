// Code generated by mockery v2.6.0. DO NOT EDIT.

package automock

import (
	servicemanager "github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/servicemanager"
	mock "github.com/stretchr/testify/mock"
)

// InstanceRemover is an autogenerated mock type for the InstanceRemover type
type InstanceRemover struct {
	mock.Mock
}

// RemoveInstance provides a mock function with given fields: smClient, instance
func (_m *InstanceRemover) RemoveInstance(smClient servicemanager.Client, instance servicemanager.InstanceKey) error {
	ret := _m.Called(smClient, instance)

	var r0 error
	if rf, ok := ret.Get(0).(func(servicemanager.Client, servicemanager.InstanceKey) error); ok {
		r0 = rf(smClient, instance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}