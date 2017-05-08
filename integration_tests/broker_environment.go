// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package integration_tests

import (
	"fmt"
	"math/rand"

	"github.com/pivotal-cf/on-demand-service-broker/config"
)

const (
	bindingId = "Gjklh45ljkhn"

	bindingPlanID    = "plan-guid-from-cc"
	bindingServiceID = "service-guid-from-cc"
	appGUID          = "app-guid-from-cc"
)

type BrokerEnvironment struct {
	Broker         *Broker
	Bosh           *Bosh
	CF             *CloudFoundry
	ServiceAdapter *ServiceAdapter
	Credhub        Credhub
}

func NewBrokerEnvironment(bosh *Bosh, cf *CloudFoundry, serviceAdapter *ServiceAdapter, credhub Credhub, brokerBinaryPath string) *BrokerEnvironment {
	return &BrokerEnvironment{
		Broker:         NewBroker(brokerBinaryPath),
		Bosh:           bosh,
		CF:             cf,
		ServiceAdapter: serviceAdapter,
		Credhub:        credhub,
	}
}

func (be *BrokerEnvironment) Start() {
	be.CF.RespondsToInitialChecks()
	be.Bosh.RespondsToInitialChecks()
	be.Broker.Start(be.Configuration())
}

func (be *BrokerEnvironment) Configuration() *config.Config {
	return &config.Config{
		Broker:         be.Broker.Configuration(),
		Bosh:           be.Bosh.Configuration(),
		CF:             be.CF.Configuration(),
		ServiceAdapter: be.ServiceAdapter.Configuration(),
		Credhub:        be.Credhub.Configuration(),
	}
}

func (be *BrokerEnvironment) Verify() {
	be.Bosh.Verify()
	be.CF.Verify()
	be.Credhub.Verify()
}

func (be *BrokerEnvironment) Close() {
	be.Broker.Close()
	be.CF.Close()
	be.Bosh.Close()
	be.Credhub.Close()
}

type ServiceInstanceID string

func AServiceInstanceID() ServiceInstanceID {
	return ServiceInstanceID(fmt.Sprintf("service-instance-ID-%d", rand.Int()))
}