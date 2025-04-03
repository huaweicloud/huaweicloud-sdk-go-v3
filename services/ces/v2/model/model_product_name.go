package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductName 产品层级跨纬规则创建时需要指明的规则产品名称，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"
type ProductName struct {
}

func (o ProductName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductName struct{}"
	}

	return strings.Join([]string{"ProductName", string(data)}, " ")
}
