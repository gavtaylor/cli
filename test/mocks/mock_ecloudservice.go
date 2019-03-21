// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ukfast/sdk-go/pkg/service/ecloud (interfaces: ECloudService)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	connection "github.com/ukfast/sdk-go/pkg/connection"
	ecloud "github.com/ukfast/sdk-go/pkg/service/ecloud"
	reflect "reflect"
)

// MockECloudService is a mock of ECloudService interface
type MockECloudService struct {
	ctrl     *gomock.Controller
	recorder *MockECloudServiceMockRecorder
}

// MockECloudServiceMockRecorder is the mock recorder for MockECloudService
type MockECloudServiceMockRecorder struct {
	mock *MockECloudService
}

// NewMockECloudService creates a new mock instance
func NewMockECloudService(ctrl *gomock.Controller) *MockECloudService {
	mock := &MockECloudService{ctrl: ctrl}
	mock.recorder = &MockECloudServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockECloudService) EXPECT() *MockECloudServiceMockRecorder {
	return m.recorder
}

// CloneVirtualMachine mocks base method
func (m *MockECloudService) CloneVirtualMachine(arg0 int, arg1 ecloud.CloneVirtualMachineRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneVirtualMachine", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneVirtualMachine indicates an expected call of CloneVirtualMachine
func (mr *MockECloudServiceMockRecorder) CloneVirtualMachine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).CloneVirtualMachine), arg0, arg1)
}

// CreateSolutionTag mocks base method
func (m *MockECloudService) CreateSolutionTag(arg0 int, arg1 ecloud.CreateTagRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSolutionTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSolutionTag indicates an expected call of CreateSolutionTag
func (mr *MockECloudServiceMockRecorder) CreateSolutionTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSolutionTag", reflect.TypeOf((*MockECloudService)(nil).CreateSolutionTag), arg0, arg1)
}

// CreateVirtualMachine mocks base method
func (m *MockECloudService) CreateVirtualMachine(arg0 ecloud.CreateVirtualMachineRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualMachine", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVirtualMachine indicates an expected call of CreateVirtualMachine
func (mr *MockECloudServiceMockRecorder) CreateVirtualMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).CreateVirtualMachine), arg0)
}

// CreateVirtualMachineTag mocks base method
func (m *MockECloudService) CreateVirtualMachineTag(arg0 int, arg1 ecloud.CreateTagRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVirtualMachineTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVirtualMachineTag indicates an expected call of CreateVirtualMachineTag
func (mr *MockECloudServiceMockRecorder) CreateVirtualMachineTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualMachineTag", reflect.TypeOf((*MockECloudService)(nil).CreateVirtualMachineTag), arg0, arg1)
}

// DeleteSolutionTag mocks base method
func (m *MockECloudService) DeleteSolutionTag(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSolutionTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSolutionTag indicates an expected call of DeleteSolutionTag
func (mr *MockECloudServiceMockRecorder) DeleteSolutionTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSolutionTag", reflect.TypeOf((*MockECloudService)(nil).DeleteSolutionTag), arg0, arg1)
}

// DeleteVirtualMachine mocks base method
func (m *MockECloudService) DeleteVirtualMachine(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVirtualMachine", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVirtualMachine indicates an expected call of DeleteVirtualMachine
func (mr *MockECloudServiceMockRecorder) DeleteVirtualMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).DeleteVirtualMachine), arg0)
}

// DeleteVirtualMachineTag mocks base method
func (m *MockECloudService) DeleteVirtualMachineTag(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVirtualMachineTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVirtualMachineTag indicates an expected call of DeleteVirtualMachineTag
func (mr *MockECloudServiceMockRecorder) DeleteVirtualMachineTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualMachineTag", reflect.TypeOf((*MockECloudService)(nil).DeleteVirtualMachineTag), arg0, arg1)
}

// GetDatastore mocks base method
func (m *MockECloudService) GetDatastore(arg0 int) (ecloud.Datastore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatastore", arg0)
	ret0, _ := ret[0].(ecloud.Datastore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatastore indicates an expected call of GetDatastore
func (mr *MockECloudServiceMockRecorder) GetDatastore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatastore", reflect.TypeOf((*MockECloudService)(nil).GetDatastore), arg0)
}

// GetDatastores mocks base method
func (m *MockECloudService) GetDatastores(arg0 connection.APIRequestParameters) ([]ecloud.Datastore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatastores", arg0)
	ret0, _ := ret[0].([]ecloud.Datastore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatastores indicates an expected call of GetDatastores
func (mr *MockECloudServiceMockRecorder) GetDatastores(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatastores", reflect.TypeOf((*MockECloudService)(nil).GetDatastores), arg0)
}

// GetDatastoresPaginated mocks base method
func (m *MockECloudService) GetDatastoresPaginated(arg0 connection.APIRequestParameters) ([]ecloud.Datastore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatastoresPaginated", arg0)
	ret0, _ := ret[0].([]ecloud.Datastore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDatastoresPaginated indicates an expected call of GetDatastoresPaginated
func (mr *MockECloudServiceMockRecorder) GetDatastoresPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatastoresPaginated", reflect.TypeOf((*MockECloudService)(nil).GetDatastoresPaginated), arg0)
}

// GetFirewall mocks base method
func (m *MockECloudService) GetFirewall(arg0 int) (ecloud.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirewall", arg0)
	ret0, _ := ret[0].(ecloud.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirewall indicates an expected call of GetFirewall
func (mr *MockECloudServiceMockRecorder) GetFirewall(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirewall", reflect.TypeOf((*MockECloudService)(nil).GetFirewall), arg0)
}

// GetFirewallConfig mocks base method
func (m *MockECloudService) GetFirewallConfig(arg0 int) (ecloud.FirewallConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirewallConfig", arg0)
	ret0, _ := ret[0].(ecloud.FirewallConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirewallConfig indicates an expected call of GetFirewallConfig
func (mr *MockECloudServiceMockRecorder) GetFirewallConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirewallConfig", reflect.TypeOf((*MockECloudService)(nil).GetFirewallConfig), arg0)
}

// GetFirewalls mocks base method
func (m *MockECloudService) GetFirewalls(arg0 connection.APIRequestParameters) ([]ecloud.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirewalls", arg0)
	ret0, _ := ret[0].([]ecloud.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirewalls indicates an expected call of GetFirewalls
func (mr *MockECloudServiceMockRecorder) GetFirewalls(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirewalls", reflect.TypeOf((*MockECloudService)(nil).GetFirewalls), arg0)
}

// GetFirewallsPaginated mocks base method
func (m *MockECloudService) GetFirewallsPaginated(arg0 connection.APIRequestParameters) ([]ecloud.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirewallsPaginated", arg0)
	ret0, _ := ret[0].([]ecloud.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirewallsPaginated indicates an expected call of GetFirewallsPaginated
func (mr *MockECloudServiceMockRecorder) GetFirewallsPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirewallsPaginated", reflect.TypeOf((*MockECloudService)(nil).GetFirewallsPaginated), arg0)
}

// GetHost mocks base method
func (m *MockECloudService) GetHost(arg0 int) (ecloud.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHost", arg0)
	ret0, _ := ret[0].(ecloud.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHost indicates an expected call of GetHost
func (mr *MockECloudServiceMockRecorder) GetHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHost", reflect.TypeOf((*MockECloudService)(nil).GetHost), arg0)
}

// GetHosts mocks base method
func (m *MockECloudService) GetHosts(arg0 connection.APIRequestParameters) ([]ecloud.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHosts", arg0)
	ret0, _ := ret[0].([]ecloud.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHosts indicates an expected call of GetHosts
func (mr *MockECloudServiceMockRecorder) GetHosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHosts", reflect.TypeOf((*MockECloudService)(nil).GetHosts), arg0)
}

// GetHostsPaginated mocks base method
func (m *MockECloudService) GetHostsPaginated(arg0 connection.APIRequestParameters) ([]ecloud.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostsPaginated", arg0)
	ret0, _ := ret[0].([]ecloud.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostsPaginated indicates an expected call of GetHostsPaginated
func (mr *MockECloudServiceMockRecorder) GetHostsPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostsPaginated", reflect.TypeOf((*MockECloudService)(nil).GetHostsPaginated), arg0)
}

// GetPod mocks base method
func (m *MockECloudService) GetPod(arg0 int) (ecloud.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPod", arg0)
	ret0, _ := ret[0].(ecloud.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPod indicates an expected call of GetPod
func (mr *MockECloudServiceMockRecorder) GetPod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPod", reflect.TypeOf((*MockECloudService)(nil).GetPod), arg0)
}

// GetPodTemplates mocks base method
func (m *MockECloudService) GetPodTemplates(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodTemplates", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodTemplates indicates an expected call of GetPodTemplates
func (mr *MockECloudServiceMockRecorder) GetPodTemplates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodTemplates", reflect.TypeOf((*MockECloudService)(nil).GetPodTemplates), arg0, arg1)
}

// GetPodTemplatesPaginated mocks base method
func (m *MockECloudService) GetPodTemplatesPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodTemplatesPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodTemplatesPaginated indicates an expected call of GetPodTemplatesPaginated
func (mr *MockECloudServiceMockRecorder) GetPodTemplatesPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodTemplatesPaginated", reflect.TypeOf((*MockECloudService)(nil).GetPodTemplatesPaginated), arg0, arg1)
}

// GetPods mocks base method
func (m *MockECloudService) GetPods(arg0 connection.APIRequestParameters) ([]ecloud.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPods", arg0)
	ret0, _ := ret[0].([]ecloud.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPods indicates an expected call of GetPods
func (mr *MockECloudServiceMockRecorder) GetPods(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPods", reflect.TypeOf((*MockECloudService)(nil).GetPods), arg0)
}

// GetPodsPaginated mocks base method
func (m *MockECloudService) GetPodsPaginated(arg0 connection.APIRequestParameters) ([]ecloud.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodsPaginated", arg0)
	ret0, _ := ret[0].([]ecloud.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodsPaginated indicates an expected call of GetPodsPaginated
func (mr *MockECloudServiceMockRecorder) GetPodsPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodsPaginated", reflect.TypeOf((*MockECloudService)(nil).GetPodsPaginated), arg0)
}

// GetSite mocks base method
func (m *MockECloudService) GetSite(arg0 int) (ecloud.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSite", arg0)
	ret0, _ := ret[0].(ecloud.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSite indicates an expected call of GetSite
func (mr *MockECloudServiceMockRecorder) GetSite(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSite", reflect.TypeOf((*MockECloudService)(nil).GetSite), arg0)
}

// GetSites mocks base method
func (m *MockECloudService) GetSites(arg0 connection.APIRequestParameters) ([]ecloud.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSites", arg0)
	ret0, _ := ret[0].([]ecloud.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSites indicates an expected call of GetSites
func (mr *MockECloudServiceMockRecorder) GetSites(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSites", reflect.TypeOf((*MockECloudService)(nil).GetSites), arg0)
}

// GetSitesPaginated mocks base method
func (m *MockECloudService) GetSitesPaginated(arg0 connection.APIRequestParameters) ([]ecloud.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSitesPaginated", arg0)
	ret0, _ := ret[0].([]ecloud.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSitesPaginated indicates an expected call of GetSitesPaginated
func (mr *MockECloudServiceMockRecorder) GetSitesPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSitesPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSitesPaginated), arg0)
}

// GetSolution mocks base method
func (m *MockECloudService) GetSolution(arg0 int) (ecloud.Solution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolution", arg0)
	ret0, _ := ret[0].(ecloud.Solution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolution indicates an expected call of GetSolution
func (mr *MockECloudServiceMockRecorder) GetSolution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolution", reflect.TypeOf((*MockECloudService)(nil).GetSolution), arg0)
}

// GetSolutionDatastores mocks base method
func (m *MockECloudService) GetSolutionDatastores(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Datastore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionDatastores", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Datastore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionDatastores indicates an expected call of GetSolutionDatastores
func (mr *MockECloudServiceMockRecorder) GetSolutionDatastores(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionDatastores", reflect.TypeOf((*MockECloudService)(nil).GetSolutionDatastores), arg0, arg1)
}

// GetSolutionDatastoresPaginated mocks base method
func (m *MockECloudService) GetSolutionDatastoresPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Datastore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionDatastoresPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Datastore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionDatastoresPaginated indicates an expected call of GetSolutionDatastoresPaginated
func (mr *MockECloudServiceMockRecorder) GetSolutionDatastoresPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionDatastoresPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSolutionDatastoresPaginated), arg0, arg1)
}

// GetSolutionFirewalls mocks base method
func (m *MockECloudService) GetSolutionFirewalls(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionFirewalls", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionFirewalls indicates an expected call of GetSolutionFirewalls
func (mr *MockECloudServiceMockRecorder) GetSolutionFirewalls(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionFirewalls", reflect.TypeOf((*MockECloudService)(nil).GetSolutionFirewalls), arg0, arg1)
}

// GetSolutionFirewallsPaginated mocks base method
func (m *MockECloudService) GetSolutionFirewallsPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Firewall, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionFirewallsPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Firewall)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionFirewallsPaginated indicates an expected call of GetSolutionFirewallsPaginated
func (mr *MockECloudServiceMockRecorder) GetSolutionFirewallsPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionFirewallsPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSolutionFirewallsPaginated), arg0, arg1)
}

// GetSolutionHosts mocks base method
func (m *MockECloudService) GetSolutionHosts(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionHosts", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionHosts indicates an expected call of GetSolutionHosts
func (mr *MockECloudServiceMockRecorder) GetSolutionHosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionHosts", reflect.TypeOf((*MockECloudService)(nil).GetSolutionHosts), arg0, arg1)
}

// GetSolutionHostsPaginated mocks base method
func (m *MockECloudService) GetSolutionHostsPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Host, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionHostsPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionHostsPaginated indicates an expected call of GetSolutionHostsPaginated
func (mr *MockECloudServiceMockRecorder) GetSolutionHostsPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionHostsPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSolutionHostsPaginated), arg0, arg1)
}

// GetSolutionNetworks mocks base method
func (m *MockECloudService) GetSolutionNetworks(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionNetworks", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionNetworks indicates an expected call of GetSolutionNetworks
func (mr *MockECloudServiceMockRecorder) GetSolutionNetworks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionNetworks", reflect.TypeOf((*MockECloudService)(nil).GetSolutionNetworks), arg0, arg1)
}

// GetSolutionNetworksPaginated mocks base method
func (m *MockECloudService) GetSolutionNetworksPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionNetworksPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionNetworksPaginated indicates an expected call of GetSolutionNetworksPaginated
func (mr *MockECloudServiceMockRecorder) GetSolutionNetworksPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionNetworksPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSolutionNetworksPaginated), arg0, arg1)
}

// GetSolutionSites mocks base method
func (m *MockECloudService) GetSolutionSites(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionSites", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionSites indicates an expected call of GetSolutionSites
func (mr *MockECloudServiceMockRecorder) GetSolutionSites(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionSites", reflect.TypeOf((*MockECloudService)(nil).GetSolutionSites), arg0, arg1)
}

// GetSolutionSitesPaginated mocks base method
func (m *MockECloudService) GetSolutionSitesPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionSitesPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionSitesPaginated indicates an expected call of GetSolutionSitesPaginated
func (mr *MockECloudServiceMockRecorder) GetSolutionSitesPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionSitesPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSolutionSitesPaginated), arg0, arg1)
}

// GetSolutionTag mocks base method
func (m *MockECloudService) GetSolutionTag(arg0 int, arg1 string) (ecloud.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionTag", arg0, arg1)
	ret0, _ := ret[0].(ecloud.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionTag indicates an expected call of GetSolutionTag
func (mr *MockECloudServiceMockRecorder) GetSolutionTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionTag", reflect.TypeOf((*MockECloudService)(nil).GetSolutionTag), arg0, arg1)
}

// GetSolutionTags mocks base method
func (m *MockECloudService) GetSolutionTags(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionTags", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionTags indicates an expected call of GetSolutionTags
func (mr *MockECloudServiceMockRecorder) GetSolutionTags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionTags", reflect.TypeOf((*MockECloudService)(nil).GetSolutionTags), arg0, arg1)
}

// GetSolutionTagsPaginated mocks base method
func (m *MockECloudService) GetSolutionTagsPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionTagsPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionTagsPaginated indicates an expected call of GetSolutionTagsPaginated
func (mr *MockECloudServiceMockRecorder) GetSolutionTagsPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionTagsPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSolutionTagsPaginated), arg0, arg1)
}

// GetSolutionTemplates mocks base method
func (m *MockECloudService) GetSolutionTemplates(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionTemplates", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionTemplates indicates an expected call of GetSolutionTemplates
func (mr *MockECloudServiceMockRecorder) GetSolutionTemplates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionTemplates", reflect.TypeOf((*MockECloudService)(nil).GetSolutionTemplates), arg0, arg1)
}

// GetSolutionTemplatesPaginated mocks base method
func (m *MockECloudService) GetSolutionTemplatesPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Template, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionTemplatesPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Template)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionTemplatesPaginated indicates an expected call of GetSolutionTemplatesPaginated
func (mr *MockECloudServiceMockRecorder) GetSolutionTemplatesPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionTemplatesPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSolutionTemplatesPaginated), arg0, arg1)
}

// GetSolutionVirtualMachines mocks base method
func (m *MockECloudService) GetSolutionVirtualMachines(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionVirtualMachines", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionVirtualMachines indicates an expected call of GetSolutionVirtualMachines
func (mr *MockECloudServiceMockRecorder) GetSolutionVirtualMachines(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionVirtualMachines", reflect.TypeOf((*MockECloudService)(nil).GetSolutionVirtualMachines), arg0, arg1)
}

// GetSolutionVirtualMachinesPaginated mocks base method
func (m *MockECloudService) GetSolutionVirtualMachinesPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionVirtualMachinesPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionVirtualMachinesPaginated indicates an expected call of GetSolutionVirtualMachinesPaginated
func (mr *MockECloudServiceMockRecorder) GetSolutionVirtualMachinesPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionVirtualMachinesPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSolutionVirtualMachinesPaginated), arg0, arg1)
}

// GetSolutions mocks base method
func (m *MockECloudService) GetSolutions(arg0 connection.APIRequestParameters) ([]ecloud.Solution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutions", arg0)
	ret0, _ := ret[0].([]ecloud.Solution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutions indicates an expected call of GetSolutions
func (mr *MockECloudServiceMockRecorder) GetSolutions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutions", reflect.TypeOf((*MockECloudService)(nil).GetSolutions), arg0)
}

// GetSolutionsPaginated mocks base method
func (m *MockECloudService) GetSolutionsPaginated(arg0 connection.APIRequestParameters) ([]ecloud.Solution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSolutionsPaginated", arg0)
	ret0, _ := ret[0].([]ecloud.Solution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSolutionsPaginated indicates an expected call of GetSolutionsPaginated
func (mr *MockECloudServiceMockRecorder) GetSolutionsPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSolutionsPaginated", reflect.TypeOf((*MockECloudService)(nil).GetSolutionsPaginated), arg0)
}

// GetVirtualMachine mocks base method
func (m *MockECloudService) GetVirtualMachine(arg0 int) (ecloud.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachine", arg0)
	ret0, _ := ret[0].(ecloud.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachine indicates an expected call of GetVirtualMachine
func (mr *MockECloudServiceMockRecorder) GetVirtualMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).GetVirtualMachine), arg0)
}

// GetVirtualMachineTag mocks base method
func (m *MockECloudService) GetVirtualMachineTag(arg0 int, arg1 string) (ecloud.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineTag", arg0, arg1)
	ret0, _ := ret[0].(ecloud.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineTag indicates an expected call of GetVirtualMachineTag
func (mr *MockECloudServiceMockRecorder) GetVirtualMachineTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineTag", reflect.TypeOf((*MockECloudService)(nil).GetVirtualMachineTag), arg0, arg1)
}

// GetVirtualMachineTags mocks base method
func (m *MockECloudService) GetVirtualMachineTags(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineTags", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineTags indicates an expected call of GetVirtualMachineTags
func (mr *MockECloudServiceMockRecorder) GetVirtualMachineTags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineTags", reflect.TypeOf((*MockECloudService)(nil).GetVirtualMachineTags), arg0, arg1)
}

// GetVirtualMachineTagsPaginated mocks base method
func (m *MockECloudService) GetVirtualMachineTagsPaginated(arg0 int, arg1 connection.APIRequestParameters) ([]ecloud.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineTagsPaginated", arg0, arg1)
	ret0, _ := ret[0].([]ecloud.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineTagsPaginated indicates an expected call of GetVirtualMachineTagsPaginated
func (mr *MockECloudServiceMockRecorder) GetVirtualMachineTagsPaginated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineTagsPaginated", reflect.TypeOf((*MockECloudService)(nil).GetVirtualMachineTagsPaginated), arg0, arg1)
}

// GetVirtualMachines mocks base method
func (m *MockECloudService) GetVirtualMachines(arg0 connection.APIRequestParameters) ([]ecloud.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachines", arg0)
	ret0, _ := ret[0].([]ecloud.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachines indicates an expected call of GetVirtualMachines
func (mr *MockECloudServiceMockRecorder) GetVirtualMachines(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachines", reflect.TypeOf((*MockECloudService)(nil).GetVirtualMachines), arg0)
}

// GetVirtualMachinesPaginated mocks base method
func (m *MockECloudService) GetVirtualMachinesPaginated(arg0 connection.APIRequestParameters) ([]ecloud.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachinesPaginated", arg0)
	ret0, _ := ret[0].([]ecloud.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachinesPaginated indicates an expected call of GetVirtualMachinesPaginated
func (mr *MockECloudServiceMockRecorder) GetVirtualMachinesPaginated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachinesPaginated", reflect.TypeOf((*MockECloudService)(nil).GetVirtualMachinesPaginated), arg0)
}

// PatchSolution mocks base method
func (m *MockECloudService) PatchSolution(arg0 int, arg1 ecloud.PatchSolutionRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchSolution", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchSolution indicates an expected call of PatchSolution
func (mr *MockECloudServiceMockRecorder) PatchSolution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSolution", reflect.TypeOf((*MockECloudService)(nil).PatchSolution), arg0, arg1)
}

// PatchSolutionTag mocks base method
func (m *MockECloudService) PatchSolutionTag(arg0 int, arg1 string, arg2 ecloud.PatchTagRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchSolutionTag", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchSolutionTag indicates an expected call of PatchSolutionTag
func (mr *MockECloudServiceMockRecorder) PatchSolutionTag(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSolutionTag", reflect.TypeOf((*MockECloudService)(nil).PatchSolutionTag), arg0, arg1, arg2)
}

// PatchVirtualMachine mocks base method
func (m *MockECloudService) PatchVirtualMachine(arg0 int, arg1 ecloud.PatchVirtualMachineRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchVirtualMachine", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchVirtualMachine indicates an expected call of PatchVirtualMachine
func (mr *MockECloudServiceMockRecorder) PatchVirtualMachine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).PatchVirtualMachine), arg0, arg1)
}

// PatchVirtualMachineTag mocks base method
func (m *MockECloudService) PatchVirtualMachineTag(arg0 int, arg1 string, arg2 ecloud.PatchTagRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchVirtualMachineTag", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchVirtualMachineTag indicates an expected call of PatchVirtualMachineTag
func (mr *MockECloudServiceMockRecorder) PatchVirtualMachineTag(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchVirtualMachineTag", reflect.TypeOf((*MockECloudService)(nil).PatchVirtualMachineTag), arg0, arg1, arg2)
}

// PowerOffVirtualMachine mocks base method
func (m *MockECloudService) PowerOffVirtualMachine(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOffVirtualMachine", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PowerOffVirtualMachine indicates an expected call of PowerOffVirtualMachine
func (mr *MockECloudServiceMockRecorder) PowerOffVirtualMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOffVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).PowerOffVirtualMachine), arg0)
}

// PowerOnVirtualMachine mocks base method
func (m *MockECloudService) PowerOnVirtualMachine(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOnVirtualMachine", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PowerOnVirtualMachine indicates an expected call of PowerOnVirtualMachine
func (mr *MockECloudServiceMockRecorder) PowerOnVirtualMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOnVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).PowerOnVirtualMachine), arg0)
}

// PowerResetVirtualMachine mocks base method
func (m *MockECloudService) PowerResetVirtualMachine(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerResetVirtualMachine", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PowerResetVirtualMachine indicates an expected call of PowerResetVirtualMachine
func (mr *MockECloudServiceMockRecorder) PowerResetVirtualMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerResetVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).PowerResetVirtualMachine), arg0)
}

// PowerRestartVirtualMachine mocks base method
func (m *MockECloudService) PowerRestartVirtualMachine(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerRestartVirtualMachine", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PowerRestartVirtualMachine indicates an expected call of PowerRestartVirtualMachine
func (mr *MockECloudServiceMockRecorder) PowerRestartVirtualMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerRestartVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).PowerRestartVirtualMachine), arg0)
}

// PowerShutdownVirtualMachine mocks base method
func (m *MockECloudService) PowerShutdownVirtualMachine(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerShutdownVirtualMachine", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PowerShutdownVirtualMachine indicates an expected call of PowerShutdownVirtualMachine
func (mr *MockECloudServiceMockRecorder) PowerShutdownVirtualMachine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerShutdownVirtualMachine", reflect.TypeOf((*MockECloudService)(nil).PowerShutdownVirtualMachine), arg0)
}