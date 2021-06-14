// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	dbconnection "github.com/kyma-project/control-plane/components/provisioners-model-migrating-job/internal/persistence/dbconnection"
	dberrors "github.com/kyma-project/control-plane/components/provisioners-model-migrating-job/internal/persistence/dberrors"

	mock "github.com/stretchr/testify/mock"
)

// ReadWriteSession is an autogenerated mock type for the ReadWriteSession type
type ReadWriteSession struct {
	mock.Mock
}

// GetProviderSpecificConfigsByProvider provides a mock function with given fields: provider
func (_m *ReadWriteSession) GetProviderSpecificConfigsByProvider(provider string) ([]dbconnection.ProviderData, dberrors.Error) {
	ret := _m.Called(provider)

	var r0 []dbconnection.ProviderData
	if rf, ok := ret.Get(0).(func(string) []dbconnection.ProviderData); ok {
		r0 = rf(provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dbconnection.ProviderData)
		}
	}

	var r1 dberrors.Error
	if rf, ok := ret.Get(1).(func(string) dberrors.Error); ok {
		r1 = rf(provider)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(dberrors.Error)
		}
	}

	return r0, r1
}

// UpdateProviderSpecificConfig provides a mock function with given fields: clusterID, providerSpecificConfig
func (_m *ReadWriteSession) UpdateProviderSpecificConfig(clusterID string, providerSpecificConfig string) dberrors.Error {
	ret := _m.Called(clusterID, providerSpecificConfig)

	var r0 dberrors.Error
	if rf, ok := ret.Get(0).(func(string, string) dberrors.Error); ok {
		r0 = rf(clusterID, providerSpecificConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dberrors.Error)
		}
	}

	return r0
}