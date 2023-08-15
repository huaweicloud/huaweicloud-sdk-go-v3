package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateRouterOption
type NeutronCreateRouterOption struct {

	// 功能说明：路由器的名称。 取值范围：0-64个字符，仅支持数字、字母、中文、_(下划线)、-（中划线）、.（点）。 约束：如果name非空，则name不能重复
	Name *string `json:"name,omitempty"`

	// 功能说明：资源的管理状态 取值范围：true、false 约束：只支持true
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ExternalGatewayInfo *ExternalGatewayInfoOption `json:"external_gateway_info,omitempty"`
}

func (o NeutronCreateRouterOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateRouterOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateRouterOption", string(data)}, " ")
}
