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

// SupportedOperation サービスが提供する操作のメタデータ
type SupportedOperation struct {
	// 操作(メソッド)の名前
	Name string

	// 操作種別、種別によりIdが必要/不要が決定される
	OperationType Operations
}

// Service 各サービスが実装すべきインターフェース
type Service interface {
	// Info サービスについての情報を返す
	Info() *Info

	// Operations サポートしている操作のメタデータ一覧
	// この要素それぞれに対しxxxWithContext()が存在することが期待される
	Operations() []SupportedOperation

	// Config コンフィグ
	Config() *Config
}

// Info サービスについての情報
type Info struct {
	// サービスの名前
	Name string

	// サービスの説明
	Description string

	// 親リソースを特定するために必要なフィールドの名前
	//
	// 例えばサーバの配下のリソースであれば"ServerId"などが指定される
	ParentKeys []string
}
