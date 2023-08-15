package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTypePrimitiveTypeHolder struct {

	// 资源的类型  以HCL格式的模板为例，resource_type 为 huaweicloud_vpc  ```hcl resource \"huaweicloud_vpc\" \"my_hello_world_vpc\" {   name = \"test_vpc\" } ```  以json格式的模板为例，resource_type 为 huaweicloud_vpc  ```json {   \"resource\": {     \"huaweicloud_vpc\": {       \"my_hello_world_vpc\": {         \"name\": \"test_vpc\"       }     }   } } ```
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o ResourceTypePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTypePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ResourceTypePrimitiveTypeHolder", string(data)}, " ")
}
