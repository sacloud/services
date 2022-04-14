// Copyright 2022 The sacloud/services Authors
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

package dummy

import (
	"github.com/sacloud/services"
	"github.com/sacloud/services/meta"
	"github.com/sacloud/services/validate"
)

type Service struct{}

var _ services.Service = (*Service)(nil)

func (s *Service) Operations() []services.SupportedOperation {
	return []services.SupportedOperation{
		{
			Name:          "Find",
			OperationType: services.OperationsList,
		},
	}
}

func (s *Service) Config() *services.Config {
	return &services.Config{
		OptionDefs: []*meta.Option{
			{Key: "option1", Values: []string{"o1", "o2"}},
			{Key: "option2", Values: []string{"o3", "o4"}},
		},
	}
}

func (s *Service) Validate(p services.Parameter) error {
	return validate.New(s).Struct(p)
}

func New() *Service {
	return &Service{}
}
