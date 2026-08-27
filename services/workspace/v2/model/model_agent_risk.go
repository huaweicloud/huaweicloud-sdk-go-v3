package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentRisk Agent 风险记录
type AgentRisk struct {

	// Agent 实例 ID
	AgentId *string `json:"agent_id,omitempty"`

	// 用户名
	Username *string `json:"username,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 风险类型
	Type *string `json:"type,omitempty"`
}

func (o AgentRisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentRisk struct{}"
	}

	return strings.Join([]string{"AgentRisk", string(data)}, " ")
}
