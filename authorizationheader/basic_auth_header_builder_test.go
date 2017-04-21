// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package authorizationheader_test

import (
	"github.com/pivotal-cf/on-demand-service-broker/authorizationheader"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BasicAuth Header Builder", func() {
	It("builds basic auth header", func() {
		authBuilder := authorizationheader.NewBasicAuthHeaderBuilder("username", "password")
		authHeader, err := authBuilder.Build(logger)
		Expect(err).NotTo(HaveOccurred())
		Expect(authHeader).To(Equal("Basic dXNlcm5hbWU6cGFzc3dvcmQ="))
	})
})