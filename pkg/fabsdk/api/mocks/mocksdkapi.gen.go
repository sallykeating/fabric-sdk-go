// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/api (interfaces: CoreProviders,SvcProviders,Providers,CoreProviderFactory,ServiceProviderFactory,OrgClientFactory,SessionClientFactory)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apiconfig "github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	apicore "github.com/hyperledger/fabric-sdk-go/api/apicore"
	apicryptosuite "github.com/hyperledger/fabric-sdk-go/api/apicryptosuite"
	apifabclient "github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	chclient "github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
	chmgmtclient "github.com/hyperledger/fabric-sdk-go/api/apitxn/chmgmtclient"
	resmgmtclient "github.com/hyperledger/fabric-sdk-go/api/apitxn/resmgmtclient"
	kvstore "github.com/hyperledger/fabric-sdk-go/api/kvstore"
	api "github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/api"
)

// MockCoreProviders is a mock of CoreProviders interface
type MockCoreProviders struct {
	ctrl     *gomock.Controller
	recorder *MockCoreProvidersMockRecorder
}

// MockCoreProvidersMockRecorder is the mock recorder for MockCoreProviders
type MockCoreProvidersMockRecorder struct {
	mock *MockCoreProviders
}

// NewMockCoreProviders creates a new mock instance
func NewMockCoreProviders(ctrl *gomock.Controller) *MockCoreProviders {
	mock := &MockCoreProviders{ctrl: ctrl}
	mock.recorder = &MockCoreProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoreProviders) EXPECT() *MockCoreProvidersMockRecorder {
	return m.recorder
}

// Config mocks base method
func (m *MockCoreProviders) Config() apiconfig.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(apiconfig.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockCoreProvidersMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockCoreProviders)(nil).Config))
}

// CryptoSuite mocks base method
func (m *MockCoreProviders) CryptoSuite() apicryptosuite.CryptoSuite {
	ret := m.ctrl.Call(m, "CryptoSuite")
	ret0, _ := ret[0].(apicryptosuite.CryptoSuite)
	return ret0
}

// CryptoSuite indicates an expected call of CryptoSuite
func (mr *MockCoreProvidersMockRecorder) CryptoSuite() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptoSuite", reflect.TypeOf((*MockCoreProviders)(nil).CryptoSuite))
}

// FabricProvider mocks base method
func (m *MockCoreProviders) FabricProvider() apicore.FabricProvider {
	ret := m.ctrl.Call(m, "FabricProvider")
	ret0, _ := ret[0].(apicore.FabricProvider)
	return ret0
}

// FabricProvider indicates an expected call of FabricProvider
func (mr *MockCoreProvidersMockRecorder) FabricProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FabricProvider", reflect.TypeOf((*MockCoreProviders)(nil).FabricProvider))
}

// SigningManager mocks base method
func (m *MockCoreProviders) SigningManager() apifabclient.SigningManager {
	ret := m.ctrl.Call(m, "SigningManager")
	ret0, _ := ret[0].(apifabclient.SigningManager)
	return ret0
}

// SigningManager indicates an expected call of SigningManager
func (mr *MockCoreProvidersMockRecorder) SigningManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigningManager", reflect.TypeOf((*MockCoreProviders)(nil).SigningManager))
}

// StateStore mocks base method
func (m *MockCoreProviders) StateStore() kvstore.KVStore {
	ret := m.ctrl.Call(m, "StateStore")
	ret0, _ := ret[0].(kvstore.KVStore)
	return ret0
}

// StateStore indicates an expected call of StateStore
func (mr *MockCoreProvidersMockRecorder) StateStore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateStore", reflect.TypeOf((*MockCoreProviders)(nil).StateStore))
}

// MockSvcProviders is a mock of SvcProviders interface
type MockSvcProviders struct {
	ctrl     *gomock.Controller
	recorder *MockSvcProvidersMockRecorder
}

// MockSvcProvidersMockRecorder is the mock recorder for MockSvcProviders
type MockSvcProvidersMockRecorder struct {
	mock *MockSvcProviders
}

// NewMockSvcProviders creates a new mock instance
func NewMockSvcProviders(ctrl *gomock.Controller) *MockSvcProviders {
	mock := &MockSvcProviders{ctrl: ctrl}
	mock.recorder = &MockSvcProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSvcProviders) EXPECT() *MockSvcProvidersMockRecorder {
	return m.recorder
}

// ChannelProvider mocks base method
func (m *MockSvcProviders) ChannelProvider() apifabclient.ChannelProvider {
	ret := m.ctrl.Call(m, "ChannelProvider")
	ret0, _ := ret[0].(apifabclient.ChannelProvider)
	return ret0
}

// ChannelProvider indicates an expected call of ChannelProvider
func (mr *MockSvcProvidersMockRecorder) ChannelProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelProvider", reflect.TypeOf((*MockSvcProviders)(nil).ChannelProvider))
}

// DiscoveryProvider mocks base method
func (m *MockSvcProviders) DiscoveryProvider() apifabclient.DiscoveryProvider {
	ret := m.ctrl.Call(m, "DiscoveryProvider")
	ret0, _ := ret[0].(apifabclient.DiscoveryProvider)
	return ret0
}

// DiscoveryProvider indicates an expected call of DiscoveryProvider
func (mr *MockSvcProvidersMockRecorder) DiscoveryProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoveryProvider", reflect.TypeOf((*MockSvcProviders)(nil).DiscoveryProvider))
}

// SelectionProvider mocks base method
func (m *MockSvcProviders) SelectionProvider() apifabclient.SelectionProvider {
	ret := m.ctrl.Call(m, "SelectionProvider")
	ret0, _ := ret[0].(apifabclient.SelectionProvider)
	return ret0
}

// SelectionProvider indicates an expected call of SelectionProvider
func (mr *MockSvcProvidersMockRecorder) SelectionProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectionProvider", reflect.TypeOf((*MockSvcProviders)(nil).SelectionProvider))
}

// MockProviders is a mock of Providers interface
type MockProviders struct {
	ctrl     *gomock.Controller
	recorder *MockProvidersMockRecorder
}

// MockProvidersMockRecorder is the mock recorder for MockProviders
type MockProvidersMockRecorder struct {
	mock *MockProviders
}

// NewMockProviders creates a new mock instance
func NewMockProviders(ctrl *gomock.Controller) *MockProviders {
	mock := &MockProviders{ctrl: ctrl}
	mock.recorder = &MockProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProviders) EXPECT() *MockProvidersMockRecorder {
	return m.recorder
}

// ChannelProvider mocks base method
func (m *MockProviders) ChannelProvider() apifabclient.ChannelProvider {
	ret := m.ctrl.Call(m, "ChannelProvider")
	ret0, _ := ret[0].(apifabclient.ChannelProvider)
	return ret0
}

// ChannelProvider indicates an expected call of ChannelProvider
func (mr *MockProvidersMockRecorder) ChannelProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelProvider", reflect.TypeOf((*MockProviders)(nil).ChannelProvider))
}

// Config mocks base method
func (m *MockProviders) Config() apiconfig.Config {
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(apiconfig.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockProvidersMockRecorder) Config() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockProviders)(nil).Config))
}

// CryptoSuite mocks base method
func (m *MockProviders) CryptoSuite() apicryptosuite.CryptoSuite {
	ret := m.ctrl.Call(m, "CryptoSuite")
	ret0, _ := ret[0].(apicryptosuite.CryptoSuite)
	return ret0
}

// CryptoSuite indicates an expected call of CryptoSuite
func (mr *MockProvidersMockRecorder) CryptoSuite() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptoSuite", reflect.TypeOf((*MockProviders)(nil).CryptoSuite))
}

// DiscoveryProvider mocks base method
func (m *MockProviders) DiscoveryProvider() apifabclient.DiscoveryProvider {
	ret := m.ctrl.Call(m, "DiscoveryProvider")
	ret0, _ := ret[0].(apifabclient.DiscoveryProvider)
	return ret0
}

// DiscoveryProvider indicates an expected call of DiscoveryProvider
func (mr *MockProvidersMockRecorder) DiscoveryProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoveryProvider", reflect.TypeOf((*MockProviders)(nil).DiscoveryProvider))
}

// FabricProvider mocks base method
func (m *MockProviders) FabricProvider() apicore.FabricProvider {
	ret := m.ctrl.Call(m, "FabricProvider")
	ret0, _ := ret[0].(apicore.FabricProvider)
	return ret0
}

// FabricProvider indicates an expected call of FabricProvider
func (mr *MockProvidersMockRecorder) FabricProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FabricProvider", reflect.TypeOf((*MockProviders)(nil).FabricProvider))
}

// SelectionProvider mocks base method
func (m *MockProviders) SelectionProvider() apifabclient.SelectionProvider {
	ret := m.ctrl.Call(m, "SelectionProvider")
	ret0, _ := ret[0].(apifabclient.SelectionProvider)
	return ret0
}

// SelectionProvider indicates an expected call of SelectionProvider
func (mr *MockProvidersMockRecorder) SelectionProvider() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectionProvider", reflect.TypeOf((*MockProviders)(nil).SelectionProvider))
}

// SigningManager mocks base method
func (m *MockProviders) SigningManager() apifabclient.SigningManager {
	ret := m.ctrl.Call(m, "SigningManager")
	ret0, _ := ret[0].(apifabclient.SigningManager)
	return ret0
}

// SigningManager indicates an expected call of SigningManager
func (mr *MockProvidersMockRecorder) SigningManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigningManager", reflect.TypeOf((*MockProviders)(nil).SigningManager))
}

// StateStore mocks base method
func (m *MockProviders) StateStore() kvstore.KVStore {
	ret := m.ctrl.Call(m, "StateStore")
	ret0, _ := ret[0].(kvstore.KVStore)
	return ret0
}

// StateStore indicates an expected call of StateStore
func (mr *MockProvidersMockRecorder) StateStore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateStore", reflect.TypeOf((*MockProviders)(nil).StateStore))
}

// MockCoreProviderFactory is a mock of CoreProviderFactory interface
type MockCoreProviderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockCoreProviderFactoryMockRecorder
}

// MockCoreProviderFactoryMockRecorder is the mock recorder for MockCoreProviderFactory
type MockCoreProviderFactoryMockRecorder struct {
	mock *MockCoreProviderFactory
}

// NewMockCoreProviderFactory creates a new mock instance
func NewMockCoreProviderFactory(ctrl *gomock.Controller) *MockCoreProviderFactory {
	mock := &MockCoreProviderFactory{ctrl: ctrl}
	mock.recorder = &MockCoreProviderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoreProviderFactory) EXPECT() *MockCoreProviderFactoryMockRecorder {
	return m.recorder
}

// NewCryptoSuiteProvider mocks base method
func (m *MockCoreProviderFactory) NewCryptoSuiteProvider(arg0 apiconfig.Config) (apicryptosuite.CryptoSuite, error) {
	ret := m.ctrl.Call(m, "NewCryptoSuiteProvider", arg0)
	ret0, _ := ret[0].(apicryptosuite.CryptoSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCryptoSuiteProvider indicates an expected call of NewCryptoSuiteProvider
func (mr *MockCoreProviderFactoryMockRecorder) NewCryptoSuiteProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCryptoSuiteProvider", reflect.TypeOf((*MockCoreProviderFactory)(nil).NewCryptoSuiteProvider), arg0)
}

// NewFabricProvider mocks base method
func (m *MockCoreProviderFactory) NewFabricProvider(arg0 apifabclient.ProviderContext) (apicore.FabricProvider, error) {
	ret := m.ctrl.Call(m, "NewFabricProvider", arg0)
	ret0, _ := ret[0].(apicore.FabricProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewFabricProvider indicates an expected call of NewFabricProvider
func (mr *MockCoreProviderFactoryMockRecorder) NewFabricProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFabricProvider", reflect.TypeOf((*MockCoreProviderFactory)(nil).NewFabricProvider), arg0)
}

// NewSigningManager mocks base method
func (m *MockCoreProviderFactory) NewSigningManager(arg0 apicryptosuite.CryptoSuite, arg1 apiconfig.Config) (apifabclient.SigningManager, error) {
	ret := m.ctrl.Call(m, "NewSigningManager", arg0, arg1)
	ret0, _ := ret[0].(apifabclient.SigningManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSigningManager indicates an expected call of NewSigningManager
func (mr *MockCoreProviderFactoryMockRecorder) NewSigningManager(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSigningManager", reflect.TypeOf((*MockCoreProviderFactory)(nil).NewSigningManager), arg0, arg1)
}

// NewStateStoreProvider mocks base method
func (m *MockCoreProviderFactory) NewStateStoreProvider(arg0 apiconfig.Config) (kvstore.KVStore, error) {
	ret := m.ctrl.Call(m, "NewStateStoreProvider", arg0)
	ret0, _ := ret[0].(kvstore.KVStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStateStoreProvider indicates an expected call of NewStateStoreProvider
func (mr *MockCoreProviderFactoryMockRecorder) NewStateStoreProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStateStoreProvider", reflect.TypeOf((*MockCoreProviderFactory)(nil).NewStateStoreProvider), arg0)
}

// MockServiceProviderFactory is a mock of ServiceProviderFactory interface
type MockServiceProviderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProviderFactoryMockRecorder
}

// MockServiceProviderFactoryMockRecorder is the mock recorder for MockServiceProviderFactory
type MockServiceProviderFactoryMockRecorder struct {
	mock *MockServiceProviderFactory
}

// NewMockServiceProviderFactory creates a new mock instance
func NewMockServiceProviderFactory(ctrl *gomock.Controller) *MockServiceProviderFactory {
	mock := &MockServiceProviderFactory{ctrl: ctrl}
	mock.recorder = &MockServiceProviderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceProviderFactory) EXPECT() *MockServiceProviderFactoryMockRecorder {
	return m.recorder
}

// NewDiscoveryProvider mocks base method
func (m *MockServiceProviderFactory) NewDiscoveryProvider(arg0 apiconfig.Config) (apifabclient.DiscoveryProvider, error) {
	ret := m.ctrl.Call(m, "NewDiscoveryProvider", arg0)
	ret0, _ := ret[0].(apifabclient.DiscoveryProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDiscoveryProvider indicates an expected call of NewDiscoveryProvider
func (mr *MockServiceProviderFactoryMockRecorder) NewDiscoveryProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDiscoveryProvider", reflect.TypeOf((*MockServiceProviderFactory)(nil).NewDiscoveryProvider), arg0)
}

// NewSelectionProvider mocks base method
func (m *MockServiceProviderFactory) NewSelectionProvider(arg0 apiconfig.Config) (apifabclient.SelectionProvider, error) {
	ret := m.ctrl.Call(m, "NewSelectionProvider", arg0)
	ret0, _ := ret[0].(apifabclient.SelectionProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSelectionProvider indicates an expected call of NewSelectionProvider
func (mr *MockServiceProviderFactoryMockRecorder) NewSelectionProvider(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSelectionProvider", reflect.TypeOf((*MockServiceProviderFactory)(nil).NewSelectionProvider), arg0)
}

// MockOrgClientFactory is a mock of OrgClientFactory interface
type MockOrgClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockOrgClientFactoryMockRecorder
}

// MockOrgClientFactoryMockRecorder is the mock recorder for MockOrgClientFactory
type MockOrgClientFactoryMockRecorder struct {
	mock *MockOrgClientFactory
}

// NewMockOrgClientFactory creates a new mock instance
func NewMockOrgClientFactory(ctrl *gomock.Controller) *MockOrgClientFactory {
	mock := &MockOrgClientFactory{ctrl: ctrl}
	mock.recorder = &MockOrgClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrgClientFactory) EXPECT() *MockOrgClientFactoryMockRecorder {
	return m.recorder
}

// NewCredentialManager mocks base method
func (m *MockOrgClientFactory) NewCredentialManager(arg0 string, arg1 apiconfig.Config, arg2 apicryptosuite.CryptoSuite) (apifabclient.CredentialManager, error) {
	ret := m.ctrl.Call(m, "NewCredentialManager", arg0, arg1, arg2)
	ret0, _ := ret[0].(apifabclient.CredentialManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewCredentialManager indicates an expected call of NewCredentialManager
func (mr *MockOrgClientFactoryMockRecorder) NewCredentialManager(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCredentialManager", reflect.TypeOf((*MockOrgClientFactory)(nil).NewCredentialManager), arg0, arg1, arg2)
}

// MockSessionClientFactory is a mock of SessionClientFactory interface
type MockSessionClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSessionClientFactoryMockRecorder
}

// MockSessionClientFactoryMockRecorder is the mock recorder for MockSessionClientFactory
type MockSessionClientFactoryMockRecorder struct {
	mock *MockSessionClientFactory
}

// NewMockSessionClientFactory creates a new mock instance
func NewMockSessionClientFactory(ctrl *gomock.Controller) *MockSessionClientFactory {
	mock := &MockSessionClientFactory{ctrl: ctrl}
	mock.recorder = &MockSessionClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionClientFactory) EXPECT() *MockSessionClientFactoryMockRecorder {
	return m.recorder
}

// NewChannelClient mocks base method
func (m *MockSessionClientFactory) NewChannelClient(arg0 api.Providers, arg1 api.SessionContext, arg2 string, arg3 apifabclient.TargetFilter) (chclient.ChannelClient, error) {
	ret := m.ctrl.Call(m, "NewChannelClient", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(chclient.ChannelClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewChannelClient indicates an expected call of NewChannelClient
func (mr *MockSessionClientFactoryMockRecorder) NewChannelClient(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChannelClient", reflect.TypeOf((*MockSessionClientFactory)(nil).NewChannelClient), arg0, arg1, arg2, arg3)
}

// NewChannelMgmtClient mocks base method
func (m *MockSessionClientFactory) NewChannelMgmtClient(arg0 api.Providers, arg1 api.SessionContext) (chmgmtclient.ChannelMgmtClient, error) {
	ret := m.ctrl.Call(m, "NewChannelMgmtClient", arg0, arg1)
	ret0, _ := ret[0].(chmgmtclient.ChannelMgmtClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewChannelMgmtClient indicates an expected call of NewChannelMgmtClient
func (mr *MockSessionClientFactoryMockRecorder) NewChannelMgmtClient(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChannelMgmtClient", reflect.TypeOf((*MockSessionClientFactory)(nil).NewChannelMgmtClient), arg0, arg1)
}

// NewResourceMgmtClient mocks base method
func (m *MockSessionClientFactory) NewResourceMgmtClient(arg0 api.Providers, arg1 api.SessionContext, arg2 resmgmtclient.TargetFilter) (resmgmtclient.ResourceMgmtClient, error) {
	ret := m.ctrl.Call(m, "NewResourceMgmtClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(resmgmtclient.ResourceMgmtClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewResourceMgmtClient indicates an expected call of NewResourceMgmtClient
func (mr *MockSessionClientFactoryMockRecorder) NewResourceMgmtClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewResourceMgmtClient", reflect.TypeOf((*MockSessionClientFactory)(nil).NewResourceMgmtClient), arg0, arg1, arg2)
}
