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

package helper

import (
	"fmt"

	"github.com/sacloud/services"
)

// PrintServiceMeta サービスメタデータを標準出力に整形して書き込む
func PrintServiceMeta(service services.Service) {
	ops, err := ServiceMeta(service)
	if err != nil {
		panic(err)
	}

	info := service.Info()
	fmt.Printf("======== Service[%s] ========\n", info.Name)
	fmt.Printf("Name: %s\n", info.Name)
	fmt.Printf("Description: %s\n", info.Description)
	fmt.Printf("ParentKeys: %s\n", info.ParentKeys)

	fmt.Println("Operations:")
	for _, op := range ops {
		fmt.Printf("\t%s(%s):\n", op.Operation.Name, op.Operation.OperationType)
		for _, field := range op.Parameters {
			fmt.Printf("\t\t%s\n", field.Name)

			{
				fmt.Println("\t\t\tStructField:")
				fmt.Printf("\t\t\t\tPkgPath: %s\n", field.StructField.PkgPath)
				fmt.Printf("\t\t\t\tTag: %s\n", field.StructField.Tag)
				fmt.Printf("\t\t\t\tType: %s\n", field.StructField.Type.String())
			}

			{
				fmt.Println("\t\t\tTag:")
				fmt.Printf("\t\t\t\tFlagName: %s\n", field.Tag.FlagName)
				fmt.Printf("\t\t\t\tFieldName: %s\n", field.Tag.FieldName)
				fmt.Printf("\t\t\t\tAliases: %s\n", field.Tag.Aliases)
				fmt.Printf("\t\t\t\tShorthand: %s\n", field.Tag.Shorthand)
				fmt.Printf("\t\t\t\tDescription: %s\n", field.Tag.Description)
				fmt.Printf("\t\t\t\tSquash: %t\n", field.Tag.Squash)
				fmt.Printf("\t\t\t\tIgnore: %t\n", field.Tag.Ignore)
				fmt.Printf("\t\t\t\tCategory: %s\n", field.Tag.Category)
				fmt.Printf("\t\t\t\tOrder: %d\n", field.Tag.Order)
				fmt.Printf("\t\t\t\tOptions: %s\n", field.Tag.Options)
				fmt.Printf("\t\t\t\tDisplayOptions: %s\n", field.Tag.DisplayOptions)
				fmt.Printf("\t\t\tLongDescription(): %#+v\n", field.Tag.LongDescription())
				fmt.Printf("\t\t\tAliasesString(): %#+v\n", field.Tag.AliasesString())
				fmt.Printf("\t\t\tOptionsString(): %#+v\n", field.Tag.OptionsString())
			}
		}
	}
}
