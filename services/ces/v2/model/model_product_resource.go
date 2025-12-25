package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductResource struct {

	// **参数解释** 资源所属的云产品名称，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\" **约束限制** 不涉及 **取值范围** 长度[0,128]个字符 **默认取值** 不涉及
	ProductName string `json:"product_name"`

	// **参数解释** 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)” **约束限制** 不涉及 **取值范围** 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值** 不涉及
	Namespace string `json:"namespace"`

	// **参数解释** 产品实例详情 **约束限制** 不涉及 **取值范围** 不超过1000个实例
	ProductInstances []ProductInstance `json:"product_instances"`
}

func (o ProductResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductResource struct{}"
	}

	return strings.Join([]string{"ProductResource", string(data)}, " ")
}
