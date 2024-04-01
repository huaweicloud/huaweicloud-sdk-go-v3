package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuerySessionResponse 会话详细信息列表。
type QuerySessionResponse struct {

	// 会话ID。
	Id string `json:"id"`

	// 当前会话是否活跃。 取值为“true”，表示活跃。 取值为“false”，表示不活跃。
	Active bool `json:"active"`

	// 操作。
	Operation string `json:"operation"`

	// 操作类型。
	Type string `json:"type"`

	// 运行时间，单位为 us。
	CostTime string `json:"cost_time"`

	// 执行计划描述。
	PlanSummary string `json:"plan_summary"`

	// 主机。
	Host string `json:"host"`

	// 客户端地址。
	Client string `json:"client"`

	// 连接描述。
	Description string `json:"description"`

	// 命名空间。
	Namespace string `json:"namespace"`
}

func (o QuerySessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySessionResponse struct{}"
	}

	return strings.Join([]string{"QuerySessionResponse", string(data)}, " ")
}
