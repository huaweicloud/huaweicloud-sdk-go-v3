package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetFirewallInstanceData struct {

	// **参数解释**： 每页显示个数 **取值范围**： 1-1024
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量：指定返回记录的开始位 **取值范围**： 大于或等于0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 防火墙总数 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 查询防火墙实例信息列表 **取值范围**： 不涉及
	Records *[]GetFirewallInstanceResponseRecord `json:"records,omitempty"`
}

func (o GetFirewallInstanceData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetFirewallInstanceData struct{}"
	}

	return strings.Join([]string{"GetFirewallInstanceData", string(data)}, " ")
}
