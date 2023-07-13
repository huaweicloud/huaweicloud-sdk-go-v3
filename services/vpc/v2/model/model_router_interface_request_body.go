package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RouterInterfaceRequestBody
type RouterInterfaceRequestBody struct {

	// 功能说明：路由器添加（或删除）接口请求参数port对应的id 约束：  - 使用端口的时候，端口上有且只有一个IP地址  - subnet_id与port_id其中之一必须指定
	PortId *string `json:"port_id,omitempty"`

	// 功能说明：路由器添加（或删除）接口请求参数subnet对应的id 约束：  - 使用子网的时候，子网上必须配置gatewayIP地址  - \"provider:network_type\"为\"geneve\"的网络不可以添加路由器  - subnet_id与port_id其中之一必须指定。
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o RouterInterfaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouterInterfaceRequestBody struct{}"
	}

	return strings.Join([]string{"RouterInterfaceRequestBody", string(data)}, " ")
}
