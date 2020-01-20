// Copyright 2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package healthserver

// MockHealthServerFactory mocks the healthServerFactory interface
type MockHealthServerFactory struct {
	serversByPort map[uint16]*mockHealthServer
}

// NewMockHealthServerFactory creates a new health server factory for testing
func NewMockHealthServerFactory() *MockHealthServerFactory {
	return &MockHealthServerFactory{
		serversByPort: map[uint16]*mockHealthServer{},
	}
}

// ServiceByPort returns the service for a given health check node port
func (m *MockHealthServerFactory) ServiceByPort(port uint16) *Service {
	if srv, ok := m.serversByPort[port]; ok {
		return srv.svc
	}
	return nil
}

func (m *MockHealthServerFactory) newHealthServer(port uint16, svc *Service) healthServer {
	m.serversByPort[port] = &mockHealthServer{
		port:    port,
		svc:     svc,
		factory: m,
	}

	return m.serversByPort[port]
}

type mockHealthServer struct {
	port    uint16
	svc     *Service
	factory *MockHealthServerFactory
}

func (m *mockHealthServer) updateService(svc *Service) {
	m.svc = svc
}

func (m *mockHealthServer) shutdown() {
	delete(m.factory.serversByPort, m.port)
}
