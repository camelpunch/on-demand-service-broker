// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/broker"
)

type FakeFeatureFlags struct {
	CFUserTriggeredUpgradesStub        func() bool
	cFUserTriggeredUpgradesMutex       sync.RWMutex
	cFUserTriggeredUpgradesArgsForCall []struct{}
	cFUserTriggeredUpgradesReturns     struct {
		result1 bool
	}
	cFUserTriggeredUpgradesReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFeatureFlags) CFUserTriggeredUpgrades() bool {
	fake.cFUserTriggeredUpgradesMutex.Lock()
	ret, specificReturn := fake.cFUserTriggeredUpgradesReturnsOnCall[len(fake.cFUserTriggeredUpgradesArgsForCall)]
	fake.cFUserTriggeredUpgradesArgsForCall = append(fake.cFUserTriggeredUpgradesArgsForCall, struct{}{})
	fake.recordInvocation("CFUserTriggeredUpgrades", []interface{}{})
	fake.cFUserTriggeredUpgradesMutex.Unlock()
	if fake.CFUserTriggeredUpgradesStub != nil {
		return fake.CFUserTriggeredUpgradesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cFUserTriggeredUpgradesReturns.result1
}

func (fake *FakeFeatureFlags) CFUserTriggeredUpgradesCallCount() int {
	fake.cFUserTriggeredUpgradesMutex.RLock()
	defer fake.cFUserTriggeredUpgradesMutex.RUnlock()
	return len(fake.cFUserTriggeredUpgradesArgsForCall)
}

func (fake *FakeFeatureFlags) CFUserTriggeredUpgradesReturns(result1 bool) {
	fake.CFUserTriggeredUpgradesStub = nil
	fake.cFUserTriggeredUpgradesReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeFeatureFlags) CFUserTriggeredUpgradesReturnsOnCall(i int, result1 bool) {
	fake.CFUserTriggeredUpgradesStub = nil
	if fake.cFUserTriggeredUpgradesReturnsOnCall == nil {
		fake.cFUserTriggeredUpgradesReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.cFUserTriggeredUpgradesReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeFeatureFlags) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cFUserTriggeredUpgradesMutex.RLock()
	defer fake.cFUserTriggeredUpgradesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFeatureFlags) recordInvocation(key string, args []interface{}) {
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

var _ broker.FeatureFlags = new(FakeFeatureFlags)