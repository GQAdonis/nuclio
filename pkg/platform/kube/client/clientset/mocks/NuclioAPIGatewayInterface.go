/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "github.com/nuclio/nuclio/pkg/platform/kube/apis/nuclio.io/v1beta1"

	watch "k8s.io/apimachinery/pkg/watch"
)

// NuclioAPIGatewayInterface is an autogenerated mock type for the NuclioAPIGatewayInterface type
type NuclioAPIGatewayInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, nuclioAPIGateway, opts
func (_m *NuclioAPIGatewayInterface) Create(ctx context.Context, nuclioAPIGateway *v1beta1.NuclioAPIGateway, opts v1.CreateOptions) (*v1beta1.NuclioAPIGateway, error) {
	ret := _m.Called(ctx, nuclioAPIGateway, opts)

	var r0 *v1beta1.NuclioAPIGateway
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.NuclioAPIGateway, v1.CreateOptions) *v1beta1.NuclioAPIGateway); ok {
		r0 = rf(ctx, nuclioAPIGateway, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.NuclioAPIGateway)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.NuclioAPIGateway, v1.CreateOptions) error); ok {
		r1 = rf(ctx, nuclioAPIGateway, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts
func (_m *NuclioAPIGatewayInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	ret := _m.Called(ctx, name, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions) error); ok {
		r0 = rf(ctx, name, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: ctx, opts, listOpts
func (_m *NuclioAPIGatewayInterface) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	ret := _m.Called(ctx, opts, listOpts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.DeleteOptions, v1.ListOptions) error); ok {
		r0 = rf(ctx, opts, listOpts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts
func (_m *NuclioAPIGatewayInterface) Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.NuclioAPIGateway, error) {
	ret := _m.Called(ctx, name, opts)

	var r0 *v1beta1.NuclioAPIGateway
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions) *v1beta1.NuclioAPIGateway); ok {
		r0 = rf(ctx, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.NuclioAPIGateway)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions) error); ok {
		r1 = rf(ctx, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, opts
func (_m *NuclioAPIGatewayInterface) List(ctx context.Context, opts v1.ListOptions) (*v1beta1.NuclioAPIGatewayList, error) {
	ret := _m.Called(ctx, opts)

	var r0 *v1beta1.NuclioAPIGatewayList
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *v1beta1.NuclioAPIGatewayList); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.NuclioAPIGatewayList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, name, pt, data, opts, subresources
func (_m *NuclioAPIGatewayInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (*v1beta1.NuclioAPIGateway, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, pt, data, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1beta1.NuclioAPIGateway
	if rf, ok := ret.Get(0).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) *v1beta1.NuclioAPIGateway); ok {
		r0 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.NuclioAPIGateway)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.PatchType, []byte, v1.PatchOptions, ...string) error); ok {
		r1 = rf(ctx, name, pt, data, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, nuclioAPIGateway, opts
func (_m *NuclioAPIGatewayInterface) Update(ctx context.Context, nuclioAPIGateway *v1beta1.NuclioAPIGateway, opts v1.UpdateOptions) (*v1beta1.NuclioAPIGateway, error) {
	ret := _m.Called(ctx, nuclioAPIGateway, opts)

	var r0 *v1beta1.NuclioAPIGateway
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta1.NuclioAPIGateway, v1.UpdateOptions) *v1beta1.NuclioAPIGateway); ok {
		r0 = rf(ctx, nuclioAPIGateway, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta1.NuclioAPIGateway)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1beta1.NuclioAPIGateway, v1.UpdateOptions) error); ok {
		r1 = rf(ctx, nuclioAPIGateway, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: ctx, opts
func (_m *NuclioAPIGatewayInterface) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(ctx, opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) watch.Interface); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}