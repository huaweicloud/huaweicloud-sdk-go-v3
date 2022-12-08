package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetFirewallInstanceData struct {

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 查询防火墙实例信息列表
	Records *[]GetFirewallInstanceResponseRecord `json:"records,omitempty"`
}

func (o GetFirewallInstanceData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetFirewallInstanceData struct{}"
	}

	return strings.Join([]string{"GetFirewallInstanceData", string(data)}, " ")
}
