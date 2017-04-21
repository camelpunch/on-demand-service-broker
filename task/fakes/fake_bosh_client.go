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
	"github.com/pivotal-cf/on-demand-service-broker/task"
)

type FakeBoshClient struct {
	DeployStub        func(manifest []byte, contextID string, logger *log.Logger) (int, error)
	deployMutex       sync.RWMutex
	deployArgsForCall []struct {
		manifest  []byte
		contextID string
		logger    *log.Logger
	}
	deployReturns struct {
		result1 int
		result2 error
	}
	deployReturnsOnCall map[int]struct {
		result1 int
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
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBoshClient) Deploy(manifest []byte, contextID string, logger *log.Logger) (int, error) {
	var manifestCopy []byte
	if manifest != nil {
		manifestCopy = make([]byte, len(manifest))
		copy(manifestCopy, manifest)
	}
	fake.deployMutex.Lock()
	ret, specificReturn := fake.deployReturnsOnCall[len(fake.deployArgsForCall)]
	fake.deployArgsForCall = append(fake.deployArgsForCall, struct {
		manifest  []byte
		contextID string
		logger    *log.Logger
	}{manifestCopy, contextID, logger})
	fake.recordInvocation("Deploy", []interface{}{manifestCopy, contextID, logger})
	fake.deployMutex.Unlock()
	if fake.DeployStub != nil {
		return fake.DeployStub(manifest, contextID, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deployReturns.result1, fake.deployReturns.result2
}

func (fake *FakeBoshClient) DeployCallCount() int {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return len(fake.deployArgsForCall)
}

func (fake *FakeBoshClient) DeployArgsForCall(i int) ([]byte, string, *log.Logger) {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return fake.deployArgsForCall[i].manifest, fake.deployArgsForCall[i].contextID, fake.deployArgsForCall[i].logger
}

func (fake *FakeBoshClient) DeployReturns(result1 int, result2 error) {
	fake.DeployStub = nil
	fake.deployReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) DeployReturnsOnCall(i int, result1 int, result2 error) {
	fake.DeployStub = nil
	if fake.deployReturnsOnCall == nil {
		fake.deployReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.deployReturnsOnCall[i] = struct {
		result1 int
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

func (fake *FakeBoshClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	fake.getTasksMutex.RLock()
	defer fake.getTasksMutex.RUnlock()
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
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

var _ task.BoshClient = new(FakeBoshClient)