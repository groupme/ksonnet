// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import categories "k8s.io/kubernetes/pkg/kubectl/categories"
import cobra "github.com/spf13/cobra"
import core "k8s.io/kubernetes/pkg/apis/core"
import discovery "k8s.io/client-go/discovery"
import internalclientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
import kubectl "k8s.io/kubernetes/pkg/kubectl"
import kubernetes "k8s.io/client-go/kubernetes"
import meta "k8s.io/apimachinery/pkg/api/meta"
import mock "github.com/stretchr/testify/mock"
import openapi "k8s.io/kubernetes/pkg/kubectl/cmd/util/openapi"
import pflag "github.com/spf13/pflag"
import plugins "k8s.io/kubernetes/pkg/kubectl/plugins"
import printers "k8s.io/kubernetes/pkg/printers"
import resource "k8s.io/kubernetes/pkg/kubectl/resource"
import rest "k8s.io/client-go/rest"
import runtime "k8s.io/apimachinery/pkg/runtime"
import schema "k8s.io/apimachinery/pkg/runtime/schema"
import time "time"

import v1 "k8s.io/api/core/v1"
import validation "k8s.io/kubernetes/pkg/kubectl/validation"

// Factory is an autogenerated mock type for the Factory type
type Factory struct {
	mock.Mock
}

// ApproximatePodTemplateForObject provides a mock function with given fields: _a0
func (_m *Factory) ApproximatePodTemplateForObject(_a0 runtime.Object) (*core.PodTemplateSpec, error) {
	ret := _m.Called(_a0)

	var r0 *core.PodTemplateSpec
	if rf, ok := ret.Get(0).(func(runtime.Object) *core.PodTemplateSpec); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.PodTemplateSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(runtime.Object) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AttachablePodForObject provides a mock function with given fields: object, timeout
func (_m *Factory) AttachablePodForObject(object runtime.Object, timeout time.Duration) (*core.Pod, error) {
	ret := _m.Called(object, timeout)

	var r0 *core.Pod
	if rf, ok := ret.Get(0).(func(runtime.Object, time.Duration) *core.Pod); ok {
		r0 = rf(object, timeout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(runtime.Object, time.Duration) error); ok {
		r1 = rf(object, timeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BareClientConfig provides a mock function with given fields:
func (_m *Factory) BareClientConfig() (*rest.Config, error) {
	ret := _m.Called()

	var r0 *rest.Config
	if rf, ok := ret.Get(0).(func() *rest.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BindExternalFlags provides a mock function with given fields: flags
func (_m *Factory) BindExternalFlags(flags *pflag.FlagSet) {
	_m.Called(flags)
}

// BindFlags provides a mock function with given fields: flags
func (_m *Factory) BindFlags(flags *pflag.FlagSet) {
	_m.Called(flags)
}

// CanBeAutoscaled provides a mock function with given fields: kind
func (_m *Factory) CanBeAutoscaled(kind schema.GroupKind) error {
	ret := _m.Called(kind)

	var r0 error
	if rf, ok := ret.Get(0).(func(schema.GroupKind) error); ok {
		r0 = rf(kind)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CanBeExposed provides a mock function with given fields: kind
func (_m *Factory) CanBeExposed(kind schema.GroupKind) error {
	ret := _m.Called(kind)

	var r0 error
	if rf, ok := ret.Get(0).(func(schema.GroupKind) error); ok {
		r0 = rf(kind)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CategoryExpander provides a mock function with given fields:
func (_m *Factory) CategoryExpander() categories.CategoryExpander {
	ret := _m.Called()

	var r0 categories.CategoryExpander
	if rf, ok := ret.Get(0).(func() categories.CategoryExpander); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(categories.CategoryExpander)
		}
	}

	return r0
}

// ClientConfig provides a mock function with given fields:
func (_m *Factory) ClientConfig() (*rest.Config, error) {
	ret := _m.Called()

	var r0 *rest.Config
	if rf, ok := ret.Get(0).(func() *rest.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Config)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientForMapping provides a mock function with given fields: mapping
func (_m *Factory) ClientForMapping(mapping *meta.RESTMapping) (resource.RESTClient, error) {
	ret := _m.Called(mapping)

	var r0 resource.RESTClient
	if rf, ok := ret.Get(0).(func(*meta.RESTMapping) resource.RESTClient); ok {
		r0 = rf(mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(resource.RESTClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*meta.RESTMapping) error); ok {
		r1 = rf(mapping)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientSet provides a mock function with given fields:
func (_m *Factory) ClientSet() (internalclientset.Interface, error) {
	ret := _m.Called()

	var r0 internalclientset.Interface
	if rf, ok := ret.Get(0).(func() internalclientset.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(internalclientset.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Command provides a mock function with given fields: cmd, showSecrets
func (_m *Factory) Command(cmd *cobra.Command, showSecrets bool) string {
	ret := _m.Called(cmd, showSecrets)

	var r0 string
	if rf, ok := ret.Get(0).(func(*cobra.Command, bool) string); ok {
		r0 = rf(cmd, showSecrets)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DefaultNamespace provides a mock function with given fields:
func (_m *Factory) DefaultNamespace() (string, bool, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DefaultResourceFilterFunc provides a mock function with given fields:
func (_m *Factory) DefaultResourceFilterFunc() kubectl.Filters {
	ret := _m.Called()

	var r0 kubectl.Filters
	if rf, ok := ret.Get(0).(func() kubectl.Filters); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kubectl.Filters)
		}
	}

	return r0
}

// Describer provides a mock function with given fields: mapping
func (_m *Factory) Describer(mapping *meta.RESTMapping) (printers.Describer, error) {
	ret := _m.Called(mapping)

	var r0 printers.Describer
	if rf, ok := ret.Get(0).(func(*meta.RESTMapping) printers.Describer); ok {
		r0 = rf(mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(printers.Describer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*meta.RESTMapping) error); ok {
		r1 = rf(mapping)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiscoveryClient provides a mock function with given fields:
func (_m *Factory) DiscoveryClient() (discovery.CachedDiscoveryInterface, error) {
	ret := _m.Called()

	var r0 discovery.CachedDiscoveryInterface
	if rf, ok := ret.Get(0).(func() discovery.CachedDiscoveryInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(discovery.CachedDiscoveryInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditorEnvs provides a mock function with given fields:
func (_m *Factory) EditorEnvs() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// FlagSet provides a mock function with given fields:
func (_m *Factory) FlagSet() *pflag.FlagSet {
	ret := _m.Called()

	var r0 *pflag.FlagSet
	if rf, ok := ret.Get(0).(func() *pflag.FlagSet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pflag.FlagSet)
		}
	}

	return r0
}

// Generators provides a mock function with given fields: cmdName
func (_m *Factory) Generators(cmdName string) map[string]kubectl.Generator {
	ret := _m.Called(cmdName)

	var r0 map[string]kubectl.Generator
	if rf, ok := ret.Get(0).(func(string) map[string]kubectl.Generator); ok {
		r0 = rf(cmdName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]kubectl.Generator)
		}
	}

	return r0
}

// HistoryViewer provides a mock function with given fields: mapping
func (_m *Factory) HistoryViewer(mapping *meta.RESTMapping) (kubectl.HistoryViewer, error) {
	ret := _m.Called(mapping)

	var r0 kubectl.HistoryViewer
	if rf, ok := ret.Get(0).(func(*meta.RESTMapping) kubectl.HistoryViewer); ok {
		r0 = rf(mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kubectl.HistoryViewer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*meta.RESTMapping) error); ok {
		r1 = rf(mapping)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KubernetesClientSet provides a mock function with given fields:
func (_m *Factory) KubernetesClientSet() (*kubernetes.Clientset, error) {
	ret := _m.Called()

	var r0 *kubernetes.Clientset
	if rf, ok := ret.Get(0).(func() *kubernetes.Clientset); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kubernetes.Clientset)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LabelsForObject provides a mock function with given fields: object
func (_m *Factory) LabelsForObject(object runtime.Object) (map[string]string, error) {
	ret := _m.Called(object)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(runtime.Object) map[string]string); ok {
		r0 = rf(object)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(runtime.Object) error); ok {
		r1 = rf(object)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogsForObject provides a mock function with given fields: object, options, timeout
func (_m *Factory) LogsForObject(object runtime.Object, options runtime.Object, timeout time.Duration) (*rest.Request, error) {
	ret := _m.Called(object, options, timeout)

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func(runtime.Object, runtime.Object, time.Duration) *rest.Request); ok {
		r0 = rf(object, options, timeout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(runtime.Object, runtime.Object, time.Duration) error); ok {
		r1 = rf(object, options, timeout)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MapBasedSelectorForObject provides a mock function with given fields: object
func (_m *Factory) MapBasedSelectorForObject(object runtime.Object) (string, error) {
	ret := _m.Called(object)

	var r0 string
	if rf, ok := ret.Get(0).(func(runtime.Object) string); ok {
		r0 = rf(object)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(runtime.Object) error); ok {
		r1 = rf(object)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBuilder provides a mock function with given fields:
func (_m *Factory) NewBuilder() *resource.Builder {
	ret := _m.Called()

	var r0 *resource.Builder
	if rf, ok := ret.Get(0).(func() *resource.Builder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resource.Builder)
		}
	}

	return r0
}

// Object provides a mock function with given fields:
func (_m *Factory) Object() (meta.RESTMapper, runtime.ObjectTyper) {
	ret := _m.Called()

	var r0 meta.RESTMapper
	if rf, ok := ret.Get(0).(func() meta.RESTMapper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(meta.RESTMapper)
		}
	}

	var r1 runtime.ObjectTyper
	if rf, ok := ret.Get(1).(func() runtime.ObjectTyper); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(runtime.ObjectTyper)
		}
	}

	return r0, r1
}

// OpenAPISchema provides a mock function with given fields:
func (_m *Factory) OpenAPISchema() (openapi.Resources, error) {
	ret := _m.Called()

	var r0 openapi.Resources
	if rf, ok := ret.Get(0).(func() openapi.Resources); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(openapi.Resources)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Pauser provides a mock function with given fields: info
func (_m *Factory) Pauser(info *resource.Info) ([]byte, error) {
	ret := _m.Called(info)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*resource.Info) []byte); ok {
		r0 = rf(info)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*resource.Info) error); ok {
		r1 = rf(info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PluginLoader provides a mock function with given fields:
func (_m *Factory) PluginLoader() plugins.PluginLoader {
	ret := _m.Called()

	var r0 plugins.PluginLoader
	if rf, ok := ret.Get(0).(func() plugins.PluginLoader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(plugins.PluginLoader)
		}
	}

	return r0
}

// PluginRunner provides a mock function with given fields:
func (_m *Factory) PluginRunner() plugins.PluginRunner {
	ret := _m.Called()

	var r0 plugins.PluginRunner
	if rf, ok := ret.Get(0).(func() plugins.PluginRunner); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(plugins.PluginRunner)
		}
	}

	return r0
}

// PortsForObject provides a mock function with given fields: object
func (_m *Factory) PortsForObject(object runtime.Object) ([]string, error) {
	ret := _m.Called(object)

	var r0 []string
	if rf, ok := ret.Get(0).(func(runtime.Object) []string); ok {
		r0 = rf(object)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(runtime.Object) error); ok {
		r1 = rf(object)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProtocolsForObject provides a mock function with given fields: object
func (_m *Factory) ProtocolsForObject(object runtime.Object) (map[string]string, error) {
	ret := _m.Called(object)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(runtime.Object) map[string]string); ok {
		r0 = rf(object)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(runtime.Object) error); ok {
		r1 = rf(object)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RESTClient provides a mock function with given fields:
func (_m *Factory) RESTClient() (*rest.RESTClient, error) {
	ret := _m.Called()

	var r0 *rest.RESTClient
	if rf, ok := ret.Get(0).(func() *rest.RESTClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.RESTClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reaper provides a mock function with given fields: mapping
func (_m *Factory) Reaper(mapping *meta.RESTMapping) (kubectl.Reaper, error) {
	ret := _m.Called(mapping)

	var r0 kubectl.Reaper
	if rf, ok := ret.Get(0).(func(*meta.RESTMapping) kubectl.Reaper); ok {
		r0 = rf(mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kubectl.Reaper)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*meta.RESTMapping) error); ok {
		r1 = rf(mapping)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveImage provides a mock function with given fields: imageName
func (_m *Factory) ResolveImage(imageName string) (string, error) {
	ret := _m.Called(imageName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(imageName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(imageName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Resumer provides a mock function with given fields: info
func (_m *Factory) Resumer(info *resource.Info) ([]byte, error) {
	ret := _m.Called(info)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(*resource.Info) []byte); ok {
		r0 = rf(info)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*resource.Info) error); ok {
		r1 = rf(info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rollbacker provides a mock function with given fields: mapping
func (_m *Factory) Rollbacker(mapping *meta.RESTMapping) (kubectl.Rollbacker, error) {
	ret := _m.Called(mapping)

	var r0 kubectl.Rollbacker
	if rf, ok := ret.Get(0).(func(*meta.RESTMapping) kubectl.Rollbacker); ok {
		r0 = rf(mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kubectl.Rollbacker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*meta.RESTMapping) error); ok {
		r1 = rf(mapping)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scaler provides a mock function with given fields: mapping
func (_m *Factory) Scaler(mapping *meta.RESTMapping) (kubectl.Scaler, error) {
	ret := _m.Called(mapping)

	var r0 kubectl.Scaler
	if rf, ok := ret.Get(0).(func(*meta.RESTMapping) kubectl.Scaler); ok {
		r0 = rf(mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kubectl.Scaler)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*meta.RESTMapping) error); ok {
		r1 = rf(mapping)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusViewer provides a mock function with given fields: mapping
func (_m *Factory) StatusViewer(mapping *meta.RESTMapping) (kubectl.StatusViewer, error) {
	ret := _m.Called(mapping)

	var r0 kubectl.StatusViewer
	if rf, ok := ret.Get(0).(func(*meta.RESTMapping) kubectl.StatusViewer); ok {
		r0 = rf(mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kubectl.StatusViewer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*meta.RESTMapping) error); ok {
		r1 = rf(mapping)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestedPodTemplateResources provides a mock function with given fields:
func (_m *Factory) SuggestedPodTemplateResources() []schema.GroupResource {
	ret := _m.Called()

	var r0 []schema.GroupResource
	if rf, ok := ret.Get(0).(func() []schema.GroupResource); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]schema.GroupResource)
		}
	}

	return r0
}

// UnstructuredClientForMapping provides a mock function with given fields: mapping
func (_m *Factory) UnstructuredClientForMapping(mapping *meta.RESTMapping) (resource.RESTClient, error) {
	ret := _m.Called(mapping)

	var r0 resource.RESTClient
	if rf, ok := ret.Get(0).(func(*meta.RESTMapping) resource.RESTClient); ok {
		r0 = rf(mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(resource.RESTClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*meta.RESTMapping) error); ok {
		r1 = rf(mapping)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePodSpecForObject provides a mock function with given fields: obj, fn
func (_m *Factory) UpdatePodSpecForObject(obj runtime.Object, fn func(*v1.PodSpec) error) (bool, error) {
	ret := _m.Called(obj, fn)

	var r0 bool
	if rf, ok := ret.Get(0).(func(runtime.Object, func(*v1.PodSpec) error) bool); ok {
		r0 = rf(obj, fn)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(runtime.Object, func(*v1.PodSpec) error) error); ok {
		r1 = rf(obj, fn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validator provides a mock function with given fields: validate
func (_m *Factory) Validator(validate bool) (validation.Schema, error) {
	ret := _m.Called(validate)

	var r0 validation.Schema
	if rf, ok := ret.Get(0).(func(bool) validation.Schema); ok {
		r0 = rf(validate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.Schema)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(validate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
