// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	types "github.com/k8snetworkplumbingwg/sriov-network-device-plugin/pkg/types"
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// NetDevice is an autogenerated mock type for the NetDevice type
type NetDevice struct {
	mock.Mock
}

// GetAPIDevice provides a mock function with given fields:
func (_m *NetDevice) GetAPIDevice() *v1beta1.Device {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAPIDevice")
	}

	var r0 *v1beta1.Device
	if rf, ok := ret.Get(0).(func() *v1beta1.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.Device)
		}
	}

	return r0
}

// GetDeviceCode provides a mock function with given fields:
func (_m *NetDevice) GetDeviceCode() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDeviceCode")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetDeviceID provides a mock function with given fields:
func (_m *NetDevice) GetDeviceID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDeviceID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetDeviceSpecs provides a mock function with given fields:
func (_m *NetDevice) GetDeviceSpecs() []*v1beta1.DeviceSpec {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDeviceSpecs")
	}

	var r0 []*v1beta1.DeviceSpec
	if rf, ok := ret.Get(0).(func() []*v1beta1.DeviceSpec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.DeviceSpec)
		}
	}

	return r0
}

// GetDriver provides a mock function with given fields:
func (_m *NetDevice) GetDriver() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetDriver")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetEnvVal provides a mock function with given fields:
func (_m *NetDevice) GetEnvVal() map[string]types.AdditionalInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetEnvVal")
	}

	var r0 map[string]types.AdditionalInfo
	if rf, ok := ret.Get(0).(func() map[string]types.AdditionalInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]types.AdditionalInfo)
		}
	}

	return r0
}

// GetFuncID provides a mock function with given fields:
func (_m *NetDevice) GetFuncID() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetFuncID")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetLinkSpeed provides a mock function with given fields:
func (_m *NetDevice) GetLinkSpeed() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLinkSpeed")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetLinkType provides a mock function with given fields:
func (_m *NetDevice) GetLinkType() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetLinkType")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetMounts provides a mock function with given fields:
func (_m *NetDevice) GetMounts() []*v1beta1.Mount {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMounts")
	}

	var r0 []*v1beta1.Mount
	if rf, ok := ret.Get(0).(func() []*v1beta1.Mount); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta1.Mount)
		}
	}

	return r0
}

// GetNetName provides a mock function with given fields:
func (_m *NetDevice) GetNetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPfNetName provides a mock function with given fields:
func (_m *NetDevice) GetPfNetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPfNetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPfPciAddr provides a mock function with given fields:
func (_m *NetDevice) GetPfPciAddr() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPfPciAddr")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetVendor provides a mock function with given fields:
func (_m *NetDevice) GetVendor() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetVendor")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsRdma provides a mock function with given fields:
func (_m *NetDevice) IsRdma() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsRdma")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewNetDevice creates a new instance of NetDevice. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNetDevice(t interface {
	mock.TestingT
	Cleanup(func())
}) *NetDevice {
	mock := &NetDevice{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
