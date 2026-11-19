package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiInstanceProtectionResultItem 修改OpenClaw防护状态结果项
type UpdateAiInstanceProtectionResultItem struct {

	// **参数解释**： 智能体ID **取值范围**： 字符长度为1-128位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 处理是否成功 **取值范围**： - true：成功 - false：失败
	Succeed *bool `json:"succeed,omitempty"`

	// **参数解释**： 错误消息，仅失败时返回 **取值范围**： 字符长度为0-256位
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o UpdateAiInstanceProtectionResultItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiInstanceProtectionResultItem struct{}"
	}

	return strings.Join([]string{"UpdateAiInstanceProtectionResultItem", string(data)}, " ")
}
