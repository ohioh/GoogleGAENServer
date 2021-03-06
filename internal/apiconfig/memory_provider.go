// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apiconfig

import (
	"context"

	"github.com/google/exposure-notifications-server/internal/model"
)

// Compile-time check to assert implementation.
var _ Provider = (*MemoryProvider)(nil)

// MemoryProvider is an Provider that stores values in-memory. It is primarily
// used for testing.
type MemoryProvider struct {
	Data map[string]*model.APIConfig
}

// NewMemoryProvider creates a new Provider that reads from a database.
func NewMemoryProvider(ctx context.Context, _ *Config) (Provider, error) {
	provider := &MemoryProvider{
		Data: make(map[string]*model.APIConfig),
	}

	return provider, nil
}

// AppConfig returns the config for the given app package name.
func (p *MemoryProvider) AppConfig(ctx context.Context, name string) (*model.APIConfig, error) {
	val, ok := p.Data[name]
	if !ok {
		return nil, AppNotFound
	}
	return val, nil
}
