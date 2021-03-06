// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeRouteRepository struct {
	ListRoutesStub        func(cb func(models.Route) bool) (apiErr error)
	listRoutesMutex       sync.RWMutex
	listRoutesArgsForCall []struct {
		cb func(models.Route) bool
	}
	listRoutesReturns struct {
		result1 error
	}
	ListAllRoutesStub        func(cb func(models.Route) bool) (apiErr error)
	listAllRoutesMutex       sync.RWMutex
	listAllRoutesArgsForCall []struct {
		cb func(models.Route) bool
	}
	listAllRoutesReturns struct {
		result1 error
	}
	FindStub        func(host string, domain models.DomainFields, path string, port int) (route models.Route, apiErr error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		host   string
		domain models.DomainFields
		path   string
		port   int
	}
	findReturns struct {
		result1 models.Route
		result2 error
	}
	CreateStub        func(host string, domain models.DomainFields, path string) (createdRoute models.Route, apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		host   string
		domain models.DomainFields
		path   string
	}
	createReturns struct {
		result1 models.Route
		result2 error
	}
	CheckIfExistsStub        func(host string, domain models.DomainFields, path string) (found bool, apiErr error)
	checkIfExistsMutex       sync.RWMutex
	checkIfExistsArgsForCall []struct {
		host   string
		domain models.DomainFields
		path   string
	}
	checkIfExistsReturns struct {
		result1 bool
		result2 error
	}
	CreateInSpaceStub        func(host, path, domainGuid, spaceGuid string, port int, randomPort bool) (createdRoute models.Route, apiErr error)
	createInSpaceMutex       sync.RWMutex
	createInSpaceArgsForCall []struct {
		host       string
		path       string
		domainGuid string
		spaceGuid  string
		port       int
		randomPort bool
	}
	createInSpaceReturns struct {
		result1 models.Route
		result2 error
	}
	BindStub        func(routeGuid, appGuid string) (apiErr error)
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		routeGuid string
		appGuid   string
	}
	bindReturns struct {
		result1 error
	}
	UnbindStub        func(routeGuid, appGuid string) (apiErr error)
	unbindMutex       sync.RWMutex
	unbindArgsForCall []struct {
		routeGuid string
		appGuid   string
	}
	unbindReturns struct {
		result1 error
	}
	DeleteStub        func(routeGuid string) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		routeGuid string
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeRouteRepository) ListRoutes(cb func(models.Route) bool) (apiErr error) {
	fake.listRoutesMutex.Lock()
	fake.listRoutesArgsForCall = append(fake.listRoutesArgsForCall, struct {
		cb func(models.Route) bool
	}{cb})
	fake.listRoutesMutex.Unlock()
	if fake.ListRoutesStub != nil {
		return fake.ListRoutesStub(cb)
	} else {
		return fake.listRoutesReturns.result1
	}
}

func (fake *FakeRouteRepository) ListRoutesCallCount() int {
	fake.listRoutesMutex.RLock()
	defer fake.listRoutesMutex.RUnlock()
	return len(fake.listRoutesArgsForCall)
}

func (fake *FakeRouteRepository) ListRoutesArgsForCall(i int) func(models.Route) bool {
	fake.listRoutesMutex.RLock()
	defer fake.listRoutesMutex.RUnlock()
	return fake.listRoutesArgsForCall[i].cb
}

func (fake *FakeRouteRepository) ListRoutesReturns(result1 error) {
	fake.ListRoutesStub = nil
	fake.listRoutesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouteRepository) ListAllRoutes(cb func(models.Route) bool) (apiErr error) {
	fake.listAllRoutesMutex.Lock()
	fake.listAllRoutesArgsForCall = append(fake.listAllRoutesArgsForCall, struct {
		cb func(models.Route) bool
	}{cb})
	fake.listAllRoutesMutex.Unlock()
	if fake.ListAllRoutesStub != nil {
		return fake.ListAllRoutesStub(cb)
	} else {
		return fake.listAllRoutesReturns.result1
	}
}

func (fake *FakeRouteRepository) ListAllRoutesCallCount() int {
	fake.listAllRoutesMutex.RLock()
	defer fake.listAllRoutesMutex.RUnlock()
	return len(fake.listAllRoutesArgsForCall)
}

func (fake *FakeRouteRepository) ListAllRoutesArgsForCall(i int) func(models.Route) bool {
	fake.listAllRoutesMutex.RLock()
	defer fake.listAllRoutesMutex.RUnlock()
	return fake.listAllRoutesArgsForCall[i].cb
}

func (fake *FakeRouteRepository) ListAllRoutesReturns(result1 error) {
	fake.ListAllRoutesStub = nil
	fake.listAllRoutesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouteRepository) Find(host string, domain models.DomainFields, path string, port int) (route models.Route, apiErr error) {
	fake.findMutex.Lock()
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		host   string
		domain models.DomainFields
		path   string
		port   int
	}{host, domain, path, port})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(host, domain, path, port)
	} else {
		return fake.findReturns.result1, fake.findReturns.result2
	}
}

func (fake *FakeRouteRepository) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeRouteRepository) FindArgsForCall(i int) (string, models.DomainFields, string, int) {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.findArgsForCall[i].host, fake.findArgsForCall[i].domain, fake.findArgsForCall[i].path, fake.findArgsForCall[i].port
}

func (fake *FakeRouteRepository) FindReturns(result1 models.Route, result2 error) {
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteRepository) Create(host string, domain models.DomainFields, path string) (createdRoute models.Route, apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		host   string
		domain models.DomainFields
		path   string
	}{host, domain, path})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(host, domain, path)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeRouteRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeRouteRepository) CreateArgsForCall(i int) (string, models.DomainFields, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].host, fake.createArgsForCall[i].domain, fake.createArgsForCall[i].path
}

func (fake *FakeRouteRepository) CreateReturns(result1 models.Route, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteRepository) CheckIfExists(host string, domain models.DomainFields, path string) (found bool, apiErr error) {
	fake.checkIfExistsMutex.Lock()
	fake.checkIfExistsArgsForCall = append(fake.checkIfExistsArgsForCall, struct {
		host   string
		domain models.DomainFields
		path   string
	}{host, domain, path})
	fake.checkIfExistsMutex.Unlock()
	if fake.CheckIfExistsStub != nil {
		return fake.CheckIfExistsStub(host, domain, path)
	} else {
		return fake.checkIfExistsReturns.result1, fake.checkIfExistsReturns.result2
	}
}

func (fake *FakeRouteRepository) CheckIfExistsCallCount() int {
	fake.checkIfExistsMutex.RLock()
	defer fake.checkIfExistsMutex.RUnlock()
	return len(fake.checkIfExistsArgsForCall)
}

func (fake *FakeRouteRepository) CheckIfExistsArgsForCall(i int) (string, models.DomainFields, string) {
	fake.checkIfExistsMutex.RLock()
	defer fake.checkIfExistsMutex.RUnlock()
	return fake.checkIfExistsArgsForCall[i].host, fake.checkIfExistsArgsForCall[i].domain, fake.checkIfExistsArgsForCall[i].path
}

func (fake *FakeRouteRepository) CheckIfExistsReturns(result1 bool, result2 error) {
	fake.CheckIfExistsStub = nil
	fake.checkIfExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteRepository) CreateInSpace(host string, path string, domainGuid string, spaceGuid string, port int, randomPort bool) (createdRoute models.Route, apiErr error) {
	fake.createInSpaceMutex.Lock()
	fake.createInSpaceArgsForCall = append(fake.createInSpaceArgsForCall, struct {
		host       string
		path       string
		domainGuid string
		spaceGuid  string
		port       int
		randomPort bool
	}{host, path, domainGuid, spaceGuid, port, randomPort})
	fake.createInSpaceMutex.Unlock()
	if fake.CreateInSpaceStub != nil {
		return fake.CreateInSpaceStub(host, path, domainGuid, spaceGuid, port, randomPort)
	} else {
		return fake.createInSpaceReturns.result1, fake.createInSpaceReturns.result2
	}
}

func (fake *FakeRouteRepository) CreateInSpaceCallCount() int {
	fake.createInSpaceMutex.RLock()
	defer fake.createInSpaceMutex.RUnlock()
	return len(fake.createInSpaceArgsForCall)
}

func (fake *FakeRouteRepository) CreateInSpaceArgsForCall(i int) (string, string, string, string, int, bool) {
	fake.createInSpaceMutex.RLock()
	defer fake.createInSpaceMutex.RUnlock()
	return fake.createInSpaceArgsForCall[i].host, fake.createInSpaceArgsForCall[i].path, fake.createInSpaceArgsForCall[i].domainGuid, fake.createInSpaceArgsForCall[i].spaceGuid, fake.createInSpaceArgsForCall[i].port, fake.createInSpaceArgsForCall[i].randomPort
}

func (fake *FakeRouteRepository) CreateInSpaceReturns(result1 models.Route, result2 error) {
	fake.CreateInSpaceStub = nil
	fake.createInSpaceReturns = struct {
		result1 models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteRepository) Bind(routeGuid string, appGuid string) (apiErr error) {
	fake.bindMutex.Lock()
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		routeGuid string
		appGuid   string
	}{routeGuid, appGuid})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(routeGuid, appGuid)
	} else {
		return fake.bindReturns.result1
	}
}

func (fake *FakeRouteRepository) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *FakeRouteRepository) BindArgsForCall(i int) (string, string) {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return fake.bindArgsForCall[i].routeGuid, fake.bindArgsForCall[i].appGuid
}

func (fake *FakeRouteRepository) BindReturns(result1 error) {
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouteRepository) Unbind(routeGuid string, appGuid string) (apiErr error) {
	fake.unbindMutex.Lock()
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		routeGuid string
		appGuid   string
	}{routeGuid, appGuid})
	fake.unbindMutex.Unlock()
	if fake.UnbindStub != nil {
		return fake.UnbindStub(routeGuid, appGuid)
	} else {
		return fake.unbindReturns.result1
	}
}

func (fake *FakeRouteRepository) UnbindCallCount() int {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}

func (fake *FakeRouteRepository) UnbindArgsForCall(i int) (string, string) {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return fake.unbindArgsForCall[i].routeGuid, fake.unbindArgsForCall[i].appGuid
}

func (fake *FakeRouteRepository) UnbindReturns(result1 error) {
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouteRepository) Delete(routeGuid string) (apiErr error) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		routeGuid string
	}{routeGuid})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(routeGuid)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeRouteRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeRouteRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].routeGuid
}

func (fake *FakeRouteRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ api.RouteRepository = new(FakeRouteRepository)
