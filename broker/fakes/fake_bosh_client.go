// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

// This file was generated by counterfeiter
package fakes

import (
	"log"
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/boshclient"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
)

type FakeBoshClient struct {
	GetTaskStub        func(taskID int, logger *log.Logger) (boshclient.BoshTask, error)
	getTaskMutex       sync.RWMutex
	getTaskArgsForCall []struct {
		taskID int
		logger *log.Logger
	}
	getTaskReturns struct {
		result1 boshclient.BoshTask
		result2 error
	}
	getTaskReturnsOnCall map[int]struct {
		result1 boshclient.BoshTask
		result2 error
	}
	GetTasksStub        func(deploymentName string, logger *log.Logger) (boshclient.BoshTasks, error)
	getTasksMutex       sync.RWMutex
	getTasksArgsForCall []struct {
		deploymentName string
		logger         *log.Logger
	}
	getTasksReturns struct {
		result1 boshclient.BoshTasks
		result2 error
	}
	getTasksReturnsOnCall map[int]struct {
		result1 boshclient.BoshTasks
		result2 error
	}
	GetNormalisedTasksByContextStub        func(deploymentName, contextID string, logger *log.Logger) (boshclient.BoshTasks, error)
	getNormalisedTasksByContextMutex       sync.RWMutex
	getNormalisedTasksByContextArgsForCall []struct {
		deploymentName string
		contextID      string
		logger         *log.Logger
	}
	getNormalisedTasksByContextReturns struct {
		result1 boshclient.BoshTasks
		result2 error
	}
	getNormalisedTasksByContextReturnsOnCall map[int]struct {
		result1 boshclient.BoshTasks
		result2 error
	}
	VMsStub        func(deploymentName string, logger *log.Logger) (bosh.BoshVMs, error)
	vMsMutex       sync.RWMutex
	vMsArgsForCall []struct {
		deploymentName string
		logger         *log.Logger
	}
	vMsReturns struct {
		result1 bosh.BoshVMs
		result2 error
	}
	vMsReturnsOnCall map[int]struct {
		result1 bosh.BoshVMs
		result2 error
	}
	GetDeploymentStub        func(name string, logger *log.Logger) ([]byte, bool, error)
	getDeploymentMutex       sync.RWMutex
	getDeploymentArgsForCall []struct {
		name   string
		logger *log.Logger
	}
	getDeploymentReturns struct {
		result1 []byte
		result2 bool
		result3 error
	}
	getDeploymentReturnsOnCall map[int]struct {
		result1 []byte
		result2 bool
		result3 error
	}
	GetDeploymentsStub        func(logger *log.Logger) ([]boshclient.BoshDeployment, error)
	getDeploymentsMutex       sync.RWMutex
	getDeploymentsArgsForCall []struct {
		logger *log.Logger
	}
	getDeploymentsReturns struct {
		result1 []boshclient.BoshDeployment
		result2 error
	}
	getDeploymentsReturnsOnCall map[int]struct {
		result1 []boshclient.BoshDeployment
		result2 error
	}
	DeleteDeploymentStub        func(name, contextID string, logger *log.Logger) (int, error)
	deleteDeploymentMutex       sync.RWMutex
	deleteDeploymentArgsForCall []struct {
		name      string
		contextID string
		logger    *log.Logger
	}
	deleteDeploymentReturns struct {
		result1 int
		result2 error
	}
	deleteDeploymentReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	GetDirectorVersionStub        func(logger *log.Logger) (boshclient.BoshDirectorVersion, error)
	getDirectorVersionMutex       sync.RWMutex
	getDirectorVersionArgsForCall []struct {
		logger *log.Logger
	}
	getDirectorVersionReturns struct {
		result1 boshclient.BoshDirectorVersion
		result2 error
	}
	getDirectorVersionReturnsOnCall map[int]struct {
		result1 boshclient.BoshDirectorVersion
		result2 error
	}
	RunErrandStub        func(deploymentName, errandName, contextID string, logger *log.Logger) (int, error)
	runErrandMutex       sync.RWMutex
	runErrandArgsForCall []struct {
		deploymentName string
		errandName     string
		contextID      string
		logger         *log.Logger
	}
	runErrandReturns struct {
		result1 int
		result2 error
	}
	runErrandReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBoshClient) GetTask(taskID int, logger *log.Logger) (boshclient.BoshTask, error) {
	fake.getTaskMutex.Lock()
	ret, specificReturn := fake.getTaskReturnsOnCall[len(fake.getTaskArgsForCall)]
	fake.getTaskArgsForCall = append(fake.getTaskArgsForCall, struct {
		taskID int
		logger *log.Logger
	}{taskID, logger})
	fake.recordInvocation("GetTask", []interface{}{taskID, logger})
	fake.getTaskMutex.Unlock()
	if fake.GetTaskStub != nil {
		return fake.GetTaskStub(taskID, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTaskReturns.result1, fake.getTaskReturns.result2
}

func (fake *FakeBoshClient) GetTaskCallCount() int {
	fake.getTaskMutex.RLock()
	defer fake.getTaskMutex.RUnlock()
	return len(fake.getTaskArgsForCall)
}

func (fake *FakeBoshClient) GetTaskArgsForCall(i int) (int, *log.Logger) {
	fake.getTaskMutex.RLock()
	defer fake.getTaskMutex.RUnlock()
	return fake.getTaskArgsForCall[i].taskID, fake.getTaskArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetTaskReturns(result1 boshclient.BoshTask, result2 error) {
	fake.GetTaskStub = nil
	fake.getTaskReturns = struct {
		result1 boshclient.BoshTask
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTaskReturnsOnCall(i int, result1 boshclient.BoshTask, result2 error) {
	fake.GetTaskStub = nil
	if fake.getTaskReturnsOnCall == nil {
		fake.getTaskReturnsOnCall = make(map[int]struct {
			result1 boshclient.BoshTask
			result2 error
		})
	}
	fake.getTaskReturnsOnCall[i] = struct {
		result1 boshclient.BoshTask
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTasks(deploymentName string, logger *log.Logger) (boshclient.BoshTasks, error) {
	fake.getTasksMutex.Lock()
	ret, specificReturn := fake.getTasksReturnsOnCall[len(fake.getTasksArgsForCall)]
	fake.getTasksArgsForCall = append(fake.getTasksArgsForCall, struct {
		deploymentName string
		logger         *log.Logger
	}{deploymentName, logger})
	fake.recordInvocation("GetTasks", []interface{}{deploymentName, logger})
	fake.getTasksMutex.Unlock()
	if fake.GetTasksStub != nil {
		return fake.GetTasksStub(deploymentName, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTasksReturns.result1, fake.getTasksReturns.result2
}

func (fake *FakeBoshClient) GetTasksCallCount() int {
	fake.getTasksMutex.RLock()
	defer fake.getTasksMutex.RUnlock()
	return len(fake.getTasksArgsForCall)
}

func (fake *FakeBoshClient) GetTasksArgsForCall(i int) (string, *log.Logger) {
	fake.getTasksMutex.RLock()
	defer fake.getTasksMutex.RUnlock()
	return fake.getTasksArgsForCall[i].deploymentName, fake.getTasksArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetTasksReturns(result1 boshclient.BoshTasks, result2 error) {
	fake.GetTasksStub = nil
	fake.getTasksReturns = struct {
		result1 boshclient.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTasksReturnsOnCall(i int, result1 boshclient.BoshTasks, result2 error) {
	fake.GetTasksStub = nil
	if fake.getTasksReturnsOnCall == nil {
		fake.getTasksReturnsOnCall = make(map[int]struct {
			result1 boshclient.BoshTasks
			result2 error
		})
	}
	fake.getTasksReturnsOnCall[i] = struct {
		result1 boshclient.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetNormalisedTasksByContext(deploymentName string, contextID string, logger *log.Logger) (boshclient.BoshTasks, error) {
	fake.getNormalisedTasksByContextMutex.Lock()
	ret, specificReturn := fake.getNormalisedTasksByContextReturnsOnCall[len(fake.getNormalisedTasksByContextArgsForCall)]
	fake.getNormalisedTasksByContextArgsForCall = append(fake.getNormalisedTasksByContextArgsForCall, struct {
		deploymentName string
		contextID      string
		logger         *log.Logger
	}{deploymentName, contextID, logger})
	fake.recordInvocation("GetNormalisedTasksByContext", []interface{}{deploymentName, contextID, logger})
	fake.getNormalisedTasksByContextMutex.Unlock()
	if fake.GetNormalisedTasksByContextStub != nil {
		return fake.GetNormalisedTasksByContextStub(deploymentName, contextID, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getNormalisedTasksByContextReturns.result1, fake.getNormalisedTasksByContextReturns.result2
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextCallCount() int {
	fake.getNormalisedTasksByContextMutex.RLock()
	defer fake.getNormalisedTasksByContextMutex.RUnlock()
	return len(fake.getNormalisedTasksByContextArgsForCall)
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextArgsForCall(i int) (string, string, *log.Logger) {
	fake.getNormalisedTasksByContextMutex.RLock()
	defer fake.getNormalisedTasksByContextMutex.RUnlock()
	return fake.getNormalisedTasksByContextArgsForCall[i].deploymentName, fake.getNormalisedTasksByContextArgsForCall[i].contextID, fake.getNormalisedTasksByContextArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextReturns(result1 boshclient.BoshTasks, result2 error) {
	fake.GetNormalisedTasksByContextStub = nil
	fake.getNormalisedTasksByContextReturns = struct {
		result1 boshclient.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextReturnsOnCall(i int, result1 boshclient.BoshTasks, result2 error) {
	fake.GetNormalisedTasksByContextStub = nil
	if fake.getNormalisedTasksByContextReturnsOnCall == nil {
		fake.getNormalisedTasksByContextReturnsOnCall = make(map[int]struct {
			result1 boshclient.BoshTasks
			result2 error
		})
	}
	fake.getNormalisedTasksByContextReturnsOnCall[i] = struct {
		result1 boshclient.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) VMs(deploymentName string, logger *log.Logger) (bosh.BoshVMs, error) {
	fake.vMsMutex.Lock()
	ret, specificReturn := fake.vMsReturnsOnCall[len(fake.vMsArgsForCall)]
	fake.vMsArgsForCall = append(fake.vMsArgsForCall, struct {
		deploymentName string
		logger         *log.Logger
	}{deploymentName, logger})
	fake.recordInvocation("VMs", []interface{}{deploymentName, logger})
	fake.vMsMutex.Unlock()
	if fake.VMsStub != nil {
		return fake.VMsStub(deploymentName, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.vMsReturns.result1, fake.vMsReturns.result2
}

func (fake *FakeBoshClient) VMsCallCount() int {
	fake.vMsMutex.RLock()
	defer fake.vMsMutex.RUnlock()
	return len(fake.vMsArgsForCall)
}

func (fake *FakeBoshClient) VMsArgsForCall(i int) (string, *log.Logger) {
	fake.vMsMutex.RLock()
	defer fake.vMsMutex.RUnlock()
	return fake.vMsArgsForCall[i].deploymentName, fake.vMsArgsForCall[i].logger
}

func (fake *FakeBoshClient) VMsReturns(result1 bosh.BoshVMs, result2 error) {
	fake.VMsStub = nil
	fake.vMsReturns = struct {
		result1 bosh.BoshVMs
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) VMsReturnsOnCall(i int, result1 bosh.BoshVMs, result2 error) {
	fake.VMsStub = nil
	if fake.vMsReturnsOnCall == nil {
		fake.vMsReturnsOnCall = make(map[int]struct {
			result1 bosh.BoshVMs
			result2 error
		})
	}
	fake.vMsReturnsOnCall[i] = struct {
		result1 bosh.BoshVMs
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetDeployment(name string, logger *log.Logger) ([]byte, bool, error) {
	fake.getDeploymentMutex.Lock()
	ret, specificReturn := fake.getDeploymentReturnsOnCall[len(fake.getDeploymentArgsForCall)]
	fake.getDeploymentArgsForCall = append(fake.getDeploymentArgsForCall, struct {
		name   string
		logger *log.Logger
	}{name, logger})
	fake.recordInvocation("GetDeployment", []interface{}{name, logger})
	fake.getDeploymentMutex.Unlock()
	if fake.GetDeploymentStub != nil {
		return fake.GetDeploymentStub(name, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getDeploymentReturns.result1, fake.getDeploymentReturns.result2, fake.getDeploymentReturns.result3
}

func (fake *FakeBoshClient) GetDeploymentCallCount() int {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return len(fake.getDeploymentArgsForCall)
}

func (fake *FakeBoshClient) GetDeploymentArgsForCall(i int) (string, *log.Logger) {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return fake.getDeploymentArgsForCall[i].name, fake.getDeploymentArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetDeploymentReturns(result1 []byte, result2 bool, result3 error) {
	fake.GetDeploymentStub = nil
	fake.getDeploymentReturns = struct {
		result1 []byte
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBoshClient) GetDeploymentReturnsOnCall(i int, result1 []byte, result2 bool, result3 error) {
	fake.GetDeploymentStub = nil
	if fake.getDeploymentReturnsOnCall == nil {
		fake.getDeploymentReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 bool
			result3 error
		})
	}
	fake.getDeploymentReturnsOnCall[i] = struct {
		result1 []byte
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBoshClient) GetDeployments(logger *log.Logger) ([]boshclient.BoshDeployment, error) {
	fake.getDeploymentsMutex.Lock()
	ret, specificReturn := fake.getDeploymentsReturnsOnCall[len(fake.getDeploymentsArgsForCall)]
	fake.getDeploymentsArgsForCall = append(fake.getDeploymentsArgsForCall, struct {
		logger *log.Logger
	}{logger})
	fake.recordInvocation("GetDeployments", []interface{}{logger})
	fake.getDeploymentsMutex.Unlock()
	if fake.GetDeploymentsStub != nil {
		return fake.GetDeploymentsStub(logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeploymentsReturns.result1, fake.getDeploymentsReturns.result2
}

func (fake *FakeBoshClient) GetDeploymentsCallCount() int {
	fake.getDeploymentsMutex.RLock()
	defer fake.getDeploymentsMutex.RUnlock()
	return len(fake.getDeploymentsArgsForCall)
}

func (fake *FakeBoshClient) GetDeploymentsArgsForCall(i int) *log.Logger {
	fake.getDeploymentsMutex.RLock()
	defer fake.getDeploymentsMutex.RUnlock()
	return fake.getDeploymentsArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetDeploymentsReturns(result1 []boshclient.BoshDeployment, result2 error) {
	fake.GetDeploymentsStub = nil
	fake.getDeploymentsReturns = struct {
		result1 []boshclient.BoshDeployment
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetDeploymentsReturnsOnCall(i int, result1 []boshclient.BoshDeployment, result2 error) {
	fake.GetDeploymentsStub = nil
	if fake.getDeploymentsReturnsOnCall == nil {
		fake.getDeploymentsReturnsOnCall = make(map[int]struct {
			result1 []boshclient.BoshDeployment
			result2 error
		})
	}
	fake.getDeploymentsReturnsOnCall[i] = struct {
		result1 []boshclient.BoshDeployment
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) DeleteDeployment(name string, contextID string, logger *log.Logger) (int, error) {
	fake.deleteDeploymentMutex.Lock()
	ret, specificReturn := fake.deleteDeploymentReturnsOnCall[len(fake.deleteDeploymentArgsForCall)]
	fake.deleteDeploymentArgsForCall = append(fake.deleteDeploymentArgsForCall, struct {
		name      string
		contextID string
		logger    *log.Logger
	}{name, contextID, logger})
	fake.recordInvocation("DeleteDeployment", []interface{}{name, contextID, logger})
	fake.deleteDeploymentMutex.Unlock()
	if fake.DeleteDeploymentStub != nil {
		return fake.DeleteDeploymentStub(name, contextID, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteDeploymentReturns.result1, fake.deleteDeploymentReturns.result2
}

func (fake *FakeBoshClient) DeleteDeploymentCallCount() int {
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	return len(fake.deleteDeploymentArgsForCall)
}

func (fake *FakeBoshClient) DeleteDeploymentArgsForCall(i int) (string, string, *log.Logger) {
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	return fake.deleteDeploymentArgsForCall[i].name, fake.deleteDeploymentArgsForCall[i].contextID, fake.deleteDeploymentArgsForCall[i].logger
}

func (fake *FakeBoshClient) DeleteDeploymentReturns(result1 int, result2 error) {
	fake.DeleteDeploymentStub = nil
	fake.deleteDeploymentReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) DeleteDeploymentReturnsOnCall(i int, result1 int, result2 error) {
	fake.DeleteDeploymentStub = nil
	if fake.deleteDeploymentReturnsOnCall == nil {
		fake.deleteDeploymentReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.deleteDeploymentReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetDirectorVersion(logger *log.Logger) (boshclient.BoshDirectorVersion, error) {
	fake.getDirectorVersionMutex.Lock()
	ret, specificReturn := fake.getDirectorVersionReturnsOnCall[len(fake.getDirectorVersionArgsForCall)]
	fake.getDirectorVersionArgsForCall = append(fake.getDirectorVersionArgsForCall, struct {
		logger *log.Logger
	}{logger})
	fake.recordInvocation("GetDirectorVersion", []interface{}{logger})
	fake.getDirectorVersionMutex.Unlock()
	if fake.GetDirectorVersionStub != nil {
		return fake.GetDirectorVersionStub(logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDirectorVersionReturns.result1, fake.getDirectorVersionReturns.result2
}

func (fake *FakeBoshClient) GetDirectorVersionCallCount() int {
	fake.getDirectorVersionMutex.RLock()
	defer fake.getDirectorVersionMutex.RUnlock()
	return len(fake.getDirectorVersionArgsForCall)
}

func (fake *FakeBoshClient) GetDirectorVersionArgsForCall(i int) *log.Logger {
	fake.getDirectorVersionMutex.RLock()
	defer fake.getDirectorVersionMutex.RUnlock()
	return fake.getDirectorVersionArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetDirectorVersionReturns(result1 boshclient.BoshDirectorVersion, result2 error) {
	fake.GetDirectorVersionStub = nil
	fake.getDirectorVersionReturns = struct {
		result1 boshclient.BoshDirectorVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetDirectorVersionReturnsOnCall(i int, result1 boshclient.BoshDirectorVersion, result2 error) {
	fake.GetDirectorVersionStub = nil
	if fake.getDirectorVersionReturnsOnCall == nil {
		fake.getDirectorVersionReturnsOnCall = make(map[int]struct {
			result1 boshclient.BoshDirectorVersion
			result2 error
		})
	}
	fake.getDirectorVersionReturnsOnCall[i] = struct {
		result1 boshclient.BoshDirectorVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) RunErrand(deploymentName string, errandName string, contextID string, logger *log.Logger) (int, error) {
	fake.runErrandMutex.Lock()
	ret, specificReturn := fake.runErrandReturnsOnCall[len(fake.runErrandArgsForCall)]
	fake.runErrandArgsForCall = append(fake.runErrandArgsForCall, struct {
		deploymentName string
		errandName     string
		contextID      string
		logger         *log.Logger
	}{deploymentName, errandName, contextID, logger})
	fake.recordInvocation("RunErrand", []interface{}{deploymentName, errandName, contextID, logger})
	fake.runErrandMutex.Unlock()
	if fake.RunErrandStub != nil {
		return fake.RunErrandStub(deploymentName, errandName, contextID, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runErrandReturns.result1, fake.runErrandReturns.result2
}

func (fake *FakeBoshClient) RunErrandCallCount() int {
	fake.runErrandMutex.RLock()
	defer fake.runErrandMutex.RUnlock()
	return len(fake.runErrandArgsForCall)
}

func (fake *FakeBoshClient) RunErrandArgsForCall(i int) (string, string, string, *log.Logger) {
	fake.runErrandMutex.RLock()
	defer fake.runErrandMutex.RUnlock()
	return fake.runErrandArgsForCall[i].deploymentName, fake.runErrandArgsForCall[i].errandName, fake.runErrandArgsForCall[i].contextID, fake.runErrandArgsForCall[i].logger
}

func (fake *FakeBoshClient) RunErrandReturns(result1 int, result2 error) {
	fake.RunErrandStub = nil
	fake.runErrandReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) RunErrandReturnsOnCall(i int, result1 int, result2 error) {
	fake.RunErrandStub = nil
	if fake.runErrandReturnsOnCall == nil {
		fake.runErrandReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.runErrandReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTaskMutex.RLock()
	defer fake.getTaskMutex.RUnlock()
	fake.getTasksMutex.RLock()
	defer fake.getTasksMutex.RUnlock()
	fake.getNormalisedTasksByContextMutex.RLock()
	defer fake.getNormalisedTasksByContextMutex.RUnlock()
	fake.vMsMutex.RLock()
	defer fake.vMsMutex.RUnlock()
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	fake.getDeploymentsMutex.RLock()
	defer fake.getDeploymentsMutex.RUnlock()
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	fake.getDirectorVersionMutex.RLock()
	defer fake.getDirectorVersionMutex.RUnlock()
	fake.runErrandMutex.RLock()
	defer fake.runErrandMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBoshClient) recordInvocation(key string, args []interface{}) {
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

var _ broker.BoshClient = new(FakeBoshClient)