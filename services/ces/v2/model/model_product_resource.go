package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductResource struct {

	// 资源所属的云产品，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"
	ProductName string `json:"product_name"`

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace string `json:"namespace"`

	// 产品实例详情
	ProductInstances []ProductInstance `json:"product_instances"`
}

func (o ProductResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductResource struct{}"
	}

	return strings.Join([]string{"ProductResource", string(data)}, " ")
}
