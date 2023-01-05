// Copyright 2022-2023 The sacloud/services Authors
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

package services

import "sort"

type Operations []SupportedOperation

// Categories 各操作のカテゴリーをアッパーキャメルケースにした上で昇順で返す
func (o *Operations) Categories() []string {
	categories := make(map[string]struct{})
	for _, op := range *o {
		categories[op.Category()] = struct{}{}
	}

	var results []string
	for k := range categories {
		results = append(results, k)
	}

	sort.Strings(results)
	return results
}
