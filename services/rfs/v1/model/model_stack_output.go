package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StackOutput 资源栈输出
type StackOutput struct {

	// 资源栈输出的名称，由用户在模板中定义  以 HCL 模板为例，name 为 vpc_id  ```hcl output \"vpc_id\" {   value = huaweicloud_vpc.my_hello_world_vpc.id } ```  以 json 模板为例，name 为 vpc_id ```json {   \"output\": {     \"vpc_id\": [       {         \"value\": \"${huaweicloud_vpc.my_hello_world_vpc.id}\"       }     ]   } } ```
	Name *string `json:"name,omitempty"`

	// 资源栈输出的描述，由用户在模板中定义
	Description *string `json:"description,omitempty"`

	// 资源栈输出的类型
	Type *string `json:"type,omitempty"`

	// 资源栈输出的值
	Value *string `json:"value,omitempty"`

	// 标识该资源栈输出是否为敏感信息，由用户在模板中定义  如果用户在模板中将该输出定义为sensitive，则返回体中该输出的value和type不会返回真实值，而是返回`<sensitive>`
	Sensitive *bool `json:"sensitive,omitempty"`
}

func (o StackOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackOutput struct{}"
	}

	return strings.Join([]string{"StackOutput", string(data)}, " ")
}
