package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源属性  以HCL格式的模板为例，资源属性的key 为 name，value 为 \"test_vpc\"  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   name = \"test_vpc\" } ```  以json格式的模板为例，资源属性的key 为 name，value 为 \"test_vpc\"  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\"       }     }   } } ```
type ResourceAttribute struct {

	// 资源属性的键
	Key *string `json:"key,omitempty"`

	// 资源属性的值
	Value *string `json:"value,omitempty"`
}

func (o ResourceAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceAttribute struct{}"
	}

	return strings.Join([]string{"ResourceAttribute", string(data)}, " ")
}
