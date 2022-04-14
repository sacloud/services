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

package services

// Operations サービスがサポートする操作の種別
type Operations int

const (
	OperationsUnknown Operations = iota
	OperationsCreate
	OperationsRead
	OperationsUpdate
	OperationsDelete
	OperationsList
	OperationsAction
)

func (o Operations) String() string {
	switch o {
	case OperationsUnknown:
		return "unknown"
	case OperationsCreate:
		return "create"
	case OperationsRead:
		return "read"
	case OperationsUpdate:
		return "update"
	case OperationsDelete:
		return "delete"
	case OperationsList:
		return "list"
	case OperationsAction:
		return "action"
	default:
		panic("got unknown Operations")
	}
}
