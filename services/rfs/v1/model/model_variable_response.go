package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VariableResponse struct {

	// 参数的名字  以HCL格式的模板为例，name 为 `my_hello_world_variable`  ```hcl variable \"my_hello_world_variable\" {   type = string   description = \"this is a variable\"   default = \"hello world\"   sensitive = false   nullable = false   validation {     condition     = length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \"hello\"     error_message = \"my_hello_world_variable should start with 'hello'.\"   } } ```  以json格式的模板为例，name 为 `my_hello_world_variable`  ```json {   \"variable\": {     \"my_hello_world_variable\": [       {         \"default\": \"hello world\",         \"description\": \"this is a variable\",         \"nullable\": false,         \"sensitive\": false,         \"type\": \"string\",         \"validation\": [           {             \"condition\": \"${length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \\\"hello\\\"}\",             \"error_message\": \"my_hello_world_variable should start with 'hello'.\"           }         ]       }     ]   } } ```
	Name *string `json:"name,omitempty"`

	// 参数的类型  以HCL格式的模板为例，type 为 `string`  ```hcl variable \"my_hello_world_variable\" {   type = string   description = \"this is a variable\"   default = \"hello world\"   sensitive = false   nullable = false   validation {     condition     = length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \"hello\"     error_message = \"my_hello_world_variable should start with 'hello'.\"   } } ```  以json格式的模板为例，type 为 `string`  ```json {   \"variable\": {     \"my_hello_world_variable\": [       {         \"default\": \"hello world\",         \"description\": \"this is a variable\",         \"nullable\": false,         \"sensitive\": false,         \"type\": \"string\",         \"validation\": [           {             \"condition\": \"${length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \\\"hello\\\"}\",             \"error_message\": \"my_hello_world_variable should start with 'hello'.\"           }         ]       }     ]   } } ```
	Type *string `json:"type,omitempty"`

	// 参数的描述  以HCL格式的模板为例，description 为 `this is a variable`  ```hcl variable \"my_hello_world_variable\" {   type = string   description = \"this is a variable\"   default = \"hello world\"   sensitive = false   nullable = false   validation {     condition     = length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \"hello\"     error_message = \"my_hello_world_variable should start with 'hello'.\"   } } ```  以json格式的模板为例，description 为 `this is a variable`  ```json {   \"variable\": {     \"my_hello_world_variable\": [       {         \"default\": \"hello world\",         \"description\": \"this is a variable\",         \"nullable\": false,         \"sensitive\": false,         \"type\": \"string\",         \"validation\": [           {             \"condition\": \"${length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \\\"hello\\\"}\",             \"error_message\": \"my_hello_world_variable should start with 'hello'.\"           }         ]       }     ]   } } ```
	Description *string `json:"description,omitempty"`

	// 参数默认值。此返回值的类型将与type保持一致  例如，对于type为string的变量，此值的返回类型为string；对于type为number的变量，此值的返回类型为number  以HCL格式的模板为例，default 为 `hello world`  ```hcl variable \"my_hello_world_variable\" {   type = string   description = \"this is a variable\"   default = \"hello world\"   sensitive = false   nullable = false   validation {     condition     = length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \"hello\"     error_message = \"my_hello_world_variable should start with 'hello'.\"   } } ```  以json格式的模板为例，default 为 `hello world`  ```json {   \"variable\": {     \"my_hello_world_variable\": [       {         \"default\": \"hello world\",         \"description\": \"this is a variable\",         \"nullable\": false,         \"sensitive\": false,         \"type\": \"string\",         \"validation\": [           {             \"condition\": \"${length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \\\"hello\\\"}\",             \"error_message\": \"my_hello_world_variable should start with 'hello'.\"           }         ]       }     ]   } } ```
	Default *interface{} `json:"default,omitempty"`

	// 参数是否为敏感字段  如果variable中没有定义sensitive，默认返回false。  以HCL格式的模板为例，sensitive 为 `false`  ```hcl variable \"my_hello_world_variable\" {   type = string   description = \"this is a variable\"   default = \"hello world\"   sensitive = false   nullable = false   validation {     condition     = length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \"hello\"     error_message = \"my_hello_world_variable should start with 'hello'.\"   } } ```  以json格式的模板为例，sensitive 为 `false`  ```json {   \"variable\": {     \"my_hello_world_variable\": [       {         \"default\": \"hello world\",         \"description\": \"this is a variable\",         \"nullable\": false,         \"sensitive\": false,         \"type\": \"string\",         \"validation\": [           {             \"condition\": \"${length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \\\"hello\\\"}\",             \"error_message\": \"my_hello_world_variable should start with 'hello'.\"           }         ]       }     ]   } } ```
	Sensitive *bool `json:"sensitive,omitempty"`

	// 参数是否可设置为null。  如果variable中没有定义nullable，默认返回true。  以HCL格式的模板为例，nullable 为 `false`  ```hcl variable \"my_hello_world_variable\" {   type = string   description = \"this is a variable\"   default = \"hello world\"   sensitive = false   nullable = false   validation {     condition     = length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \"hello\"     error_message = \"my_hello_world_variable should start with 'hello'.\"   } } ```  以json格式的模板为例，nullable 为 `false`  ```json {   \"variable\": {     \"my_hello_world_variable\": [       {         \"default\": \"hello world\",         \"description\": \"this is a variable\",         \"nullable\": false,         \"sensitive\": false,         \"type\": \"string\",         \"validation\": [           {             \"condition\": \"${length(var.my_hello_world_variable) > 0 && substr(var.my_hello_world_variable, 0, 5) == \\\"hello\\\"}\",             \"error_message\": \"my_hello_world_variable should start with 'hello'.\"           }         ]       }     ]   } } ```
	Nullable *bool `json:"nullable,omitempty"`

	// 参数的校验模块
	Validations *[]VariableValidationResponse `json:"validations,omitempty"`
}

func (o VariableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VariableResponse struct{}"
	}

	return strings.Join([]string{"VariableResponse", string(data)}, " ")
}
