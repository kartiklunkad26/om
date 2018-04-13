// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type StagedConfigService struct {
	FindStub        func(product string) (api.StagedProductsFindOutput, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		product string
	}
	findReturns struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	findReturnsOnCall map[int]struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	JobsStub        func(productGUID string) (map[string]string, error)
	jobsMutex       sync.RWMutex
	jobsArgsForCall []struct {
		productGUID string
	}
	jobsReturns struct {
		result1 map[string]string
		result2 error
	}
	jobsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	GetExistingJobConfigStub        func(productGUID, jobGUID string) (api.JobProperties, error)
	getExistingJobConfigMutex       sync.RWMutex
	getExistingJobConfigArgsForCall []struct {
		productGUID string
		jobGUID     string
	}
	getExistingJobConfigReturns struct {
		result1 api.JobProperties
		result2 error
	}
	getExistingJobConfigReturnsOnCall map[int]struct {
		result1 api.JobProperties
		result2 error
	}
	PropertiesStub        func(product string) (map[string]api.ResponseProperty, error)
	propertiesMutex       sync.RWMutex
	propertiesArgsForCall []struct {
		product string
	}
	propertiesReturns struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	propertiesReturnsOnCall map[int]struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	NetworksAndAZsStub        func(product string) (map[string]interface{}, error)
	networksAndAZsMutex       sync.RWMutex
	networksAndAZsArgsForCall []struct {
		product string
	}
	networksAndAZsReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	networksAndAZsReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StagedConfigService) Find(product string) (api.StagedProductsFindOutput, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("Find", []interface{}{product})
	fake.findMutex.Unlock()
	if fake.FindStub != nil {
		return fake.FindStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findReturns.result1, fake.findReturns.result2
}

func (fake *StagedConfigService) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *StagedConfigService) FindArgsForCall(i int) string {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return fake.findArgsForCall[i].product
}

func (fake *StagedConfigService) FindReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) FindReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsFindOutput
			result2 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) Jobs(productGUID string) (map[string]string, error) {
	fake.jobsMutex.Lock()
	ret, specificReturn := fake.jobsReturnsOnCall[len(fake.jobsArgsForCall)]
	fake.jobsArgsForCall = append(fake.jobsArgsForCall, struct {
		productGUID string
	}{productGUID})
	fake.recordInvocation("Jobs", []interface{}{productGUID})
	fake.jobsMutex.Unlock()
	if fake.JobsStub != nil {
		return fake.JobsStub(productGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.jobsReturns.result1, fake.jobsReturns.result2
}

func (fake *StagedConfigService) JobsCallCount() int {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return len(fake.jobsArgsForCall)
}

func (fake *StagedConfigService) JobsArgsForCall(i int) string {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return fake.jobsArgsForCall[i].productGUID
}

func (fake *StagedConfigService) JobsReturns(result1 map[string]string, result2 error) {
	fake.JobsStub = nil
	fake.jobsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) JobsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.JobsStub = nil
	if fake.jobsReturnsOnCall == nil {
		fake.jobsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.jobsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetExistingJobConfig(productGUID string, jobGUID string) (api.JobProperties, error) {
	fake.getExistingJobConfigMutex.Lock()
	ret, specificReturn := fake.getExistingJobConfigReturnsOnCall[len(fake.getExistingJobConfigArgsForCall)]
	fake.getExistingJobConfigArgsForCall = append(fake.getExistingJobConfigArgsForCall, struct {
		productGUID string
		jobGUID     string
	}{productGUID, jobGUID})
	fake.recordInvocation("GetExistingJobConfig", []interface{}{productGUID, jobGUID})
	fake.getExistingJobConfigMutex.Unlock()
	if fake.GetExistingJobConfigStub != nil {
		return fake.GetExistingJobConfigStub(productGUID, jobGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getExistingJobConfigReturns.result1, fake.getExistingJobConfigReturns.result2
}

func (fake *StagedConfigService) GetExistingJobConfigCallCount() int {
	fake.getExistingJobConfigMutex.RLock()
	defer fake.getExistingJobConfigMutex.RUnlock()
	return len(fake.getExistingJobConfigArgsForCall)
}

func (fake *StagedConfigService) GetExistingJobConfigArgsForCall(i int) (string, string) {
	fake.getExistingJobConfigMutex.RLock()
	defer fake.getExistingJobConfigMutex.RUnlock()
	return fake.getExistingJobConfigArgsForCall[i].productGUID, fake.getExistingJobConfigArgsForCall[i].jobGUID
}

func (fake *StagedConfigService) GetExistingJobConfigReturns(result1 api.JobProperties, result2 error) {
	fake.GetExistingJobConfigStub = nil
	fake.getExistingJobConfigReturns = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetExistingJobConfigReturnsOnCall(i int, result1 api.JobProperties, result2 error) {
	fake.GetExistingJobConfigStub = nil
	if fake.getExistingJobConfigReturnsOnCall == nil {
		fake.getExistingJobConfigReturnsOnCall = make(map[int]struct {
			result1 api.JobProperties
			result2 error
		})
	}
	fake.getExistingJobConfigReturnsOnCall[i] = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) Properties(product string) (map[string]api.ResponseProperty, error) {
	fake.propertiesMutex.Lock()
	ret, specificReturn := fake.propertiesReturnsOnCall[len(fake.propertiesArgsForCall)]
	fake.propertiesArgsForCall = append(fake.propertiesArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("Properties", []interface{}{product})
	fake.propertiesMutex.Unlock()
	if fake.PropertiesStub != nil {
		return fake.PropertiesStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.propertiesReturns.result1, fake.propertiesReturns.result2
}

func (fake *StagedConfigService) PropertiesCallCount() int {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return len(fake.propertiesArgsForCall)
}

func (fake *StagedConfigService) PropertiesArgsForCall(i int) string {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return fake.propertiesArgsForCall[i].product
}

func (fake *StagedConfigService) PropertiesReturns(result1 map[string]api.ResponseProperty, result2 error) {
	fake.PropertiesStub = nil
	fake.propertiesReturns = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) PropertiesReturnsOnCall(i int, result1 map[string]api.ResponseProperty, result2 error) {
	fake.PropertiesStub = nil
	if fake.propertiesReturnsOnCall == nil {
		fake.propertiesReturnsOnCall = make(map[int]struct {
			result1 map[string]api.ResponseProperty
			result2 error
		})
	}
	fake.propertiesReturnsOnCall[i] = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) NetworksAndAZs(product string) (map[string]interface{}, error) {
	fake.networksAndAZsMutex.Lock()
	ret, specificReturn := fake.networksAndAZsReturnsOnCall[len(fake.networksAndAZsArgsForCall)]
	fake.networksAndAZsArgsForCall = append(fake.networksAndAZsArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("NetworksAndAZs", []interface{}{product})
	fake.networksAndAZsMutex.Unlock()
	if fake.NetworksAndAZsStub != nil {
		return fake.NetworksAndAZsStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.networksAndAZsReturns.result1, fake.networksAndAZsReturns.result2
}

func (fake *StagedConfigService) NetworksAndAZsCallCount() int {
	fake.networksAndAZsMutex.RLock()
	defer fake.networksAndAZsMutex.RUnlock()
	return len(fake.networksAndAZsArgsForCall)
}

func (fake *StagedConfigService) NetworksAndAZsArgsForCall(i int) string {
	fake.networksAndAZsMutex.RLock()
	defer fake.networksAndAZsMutex.RUnlock()
	return fake.networksAndAZsArgsForCall[i].product
}

func (fake *StagedConfigService) NetworksAndAZsReturns(result1 map[string]interface{}, result2 error) {
	fake.NetworksAndAZsStub = nil
	fake.networksAndAZsReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) NetworksAndAZsReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.NetworksAndAZsStub = nil
	if fake.networksAndAZsReturnsOnCall == nil {
		fake.networksAndAZsReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.networksAndAZsReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	fake.getExistingJobConfigMutex.RLock()
	defer fake.getExistingJobConfigMutex.RUnlock()
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	fake.networksAndAZsMutex.RLock()
	defer fake.networksAndAZsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StagedConfigService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}