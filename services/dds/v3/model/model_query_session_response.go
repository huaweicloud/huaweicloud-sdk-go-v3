package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 会话详细信息列表。
type QuerySessionResponse struct {

	// 会话ID。
	Id string `json:"id" xml:"id"`

	// 当前会话是否活跃。 取值为“true”，表示活跃。 取值为“false”，表示不活跃。
	Active bool `json:"active" xml:"active"`

	// 操作。
	Operation string `json:"operation" xml:"operation"`

	// 操作类型。
	Type string `json:"type" xml:"type"`

	// 运行时间，单位为 ms。
	CostTime string `json:"cost_time" xml:"cost_time"`

	// 执行计划描述。
	PlanSummary string `json:"plan_summary" xml:"plan_summary"`

	// 主机。
	Host string `json:"host" xml:"host"`

	// 客户端地址。
	Client string `json:"client" xml:"client"`

	// 连接描述。
	Description string `json:"description" xml:"description"`

	// 命名空间。
	Namespace string `json:"namespace" xml:"namespace"`
}

func (o QuerySessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySessionResponse struct{}"
	}

	return strings.Join([]string{"QuerySessionResponse", string(data)}, " ")
}
