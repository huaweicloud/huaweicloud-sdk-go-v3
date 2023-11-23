package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEastWestFirewallRequestBody 创建东西向防火墙body体
type CreateEastWestFirewallRequestBody struct {

	// 出方向关联ER实例id
	ErId *string `json:"er_id,omitempty"`

	// inspection cidr
	InspectionCidr string `json:"inspection_cidr"`

	// 东西向防火墙模式，填写er
	Mode string `json:"mode"`
}

func (o CreateEastWestFirewallRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEastWestFirewallRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEastWestFirewallRequestBody", string(data)}, " ")
}
