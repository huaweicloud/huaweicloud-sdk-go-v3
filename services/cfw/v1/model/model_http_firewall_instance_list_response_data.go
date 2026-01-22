package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpFirewallInstanceListResponseData struct {

	// **参数解释**： 每页显示个数 **约束限制**： 不涉及 **取值范围**： 1-1024
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量：指定返回记录的开始位置 **约束限制**： 不涉及 **取值范围**： 大于或等于0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 租户项目ID **约束限制**： 不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 防火墙总数量 **约束限制**： 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 查询防火墙列表记录 **约束限制**： 不涉及
	Records *[]FirewallInstanceVo `json:"records,omitempty"`
}

func (o HttpFirewallInstanceListResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpFirewallInstanceListResponseData struct{}"
	}

	return strings.Join([]string{"HttpFirewallInstanceListResponseData", string(data)}, " ")
}
