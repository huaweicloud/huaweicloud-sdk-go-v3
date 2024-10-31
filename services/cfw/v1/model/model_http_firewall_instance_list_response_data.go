package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpFirewallInstanceListResponseData struct {

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 租户项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 防火墙总数量
	Total *int32 `json:"total,omitempty"`

	// 查询防火墙列表记录
	Records *[]FirewallInstanceVo `json:"records,omitempty"`
}

func (o HttpFirewallInstanceListResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpFirewallInstanceListResponseData struct{}"
	}

	return strings.Join([]string{"HttpFirewallInstanceListResponseData", string(data)}, " ")
}
