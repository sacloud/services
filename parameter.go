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

import (
	"fmt"
	"reflect"

	"github.com/sacloud/services/meta"
)

// Parameter 各パラメータが実装すべきインターフェース
type Parameter interface {
	// Initialize パラメータの初期値投入
	Initialize()

	// KeyFieldNames キーとなる項目のフィールド名リスト
	KeyFieldNames() []string
}

// ParameterValidator パラメータが独自のバリデーションを実装する場合に実装すべきインターフェース
type ParameterValidator interface {
	Validate() error
}

// ParameterMeta 指定のfuncのパラメータが持つ各フィールドのメタデータを取得
func ParameterMeta(service Service, funcName string) ([]meta.StructField, error) {
	parser := metaParser(service)
	instance, err := NewParameter(service, funcName)
	if err != nil {
		return nil, err
	}
	return parser.Parse(instance)
}

// NewParameter 指定のfuncのパラメータを新規作成&初期化して返す
func NewParameter(service Service, funcName string) (interface{}, error) {
	method, found := reflect.TypeOf(service).MethodByName(funcName)
	if !found {
		return nil, fmt.Errorf("method %q not found", funcName)
	}
	instance := reflect.New(method.Type.In(1).Elem()).Interface()
	if v, ok := instance.(Parameter); ok {
		v.Initialize()
	}
	return instance, nil
}

// ValidateParameter serviceのコンフィグを反映したバリデーターを用いたバリデーション
func ValidateParameter(service Service, parameter Parameter) error {
	if err := service.Validate(parameter); err != nil {
		return err
	}
	if p, ok := parameter.(ParameterValidator); ok {
		return p.Validate()
	}
	return nil
}

func metaParser(service Service) *meta.Parser {
	config := service.Config()
	tagName := config.MetaTagName
	if tagName == "" {
		tagName = meta.DefaultTagName
	}
	return &meta.Parser{
		Config: &meta.ParserConfig{
			TagName: tagName,
			Options: config.OptionDefs,
		},
	}
}
