// Copyright 2022-2025 The sacloud/services Authors
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
	"context"
	"fmt"
)

func (s *Service) ErrorRead(req *ErrorReadRequest) (*ErrorReadResult, error) {
	return s.ErrorReadWithContext(context.Background(), req)
}

func (s *Service) ErrorReadWithContext(ctx context.Context, req *ErrorReadRequest) (*ErrorReadResult, error) {
	return nil, fmt.Errorf("dummy")
}

type ErrorReadRequest struct {
}

type ErrorReadResult struct {
}
