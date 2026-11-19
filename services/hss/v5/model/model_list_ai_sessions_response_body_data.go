package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiSessionsResponseBodyData AI会话响应数据项
type ListAiSessionsResponseBodyData struct {

	// **参数解释**： 会话ID **取值范围**： 字符长度为1-128位
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**： 智能体名字 **取值范围**： 字符长度为1-128位
	AgentName *string `json:"agent_name,omitempty"`

	// **参数解释**： 智能体ID **取值范围**： 字符长度为1-128位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 智能体运行时名字 **取值范围**： 字符长度为1-128位
	RuntimeName *string `json:"runtime_name,omitempty"`

	// **参数解释**： 智能体运行时ID **取值范围**： 字符长度为1-128位
	RuntimeId *string `json:"runtime_id,omitempty"`

	// **参数解释**： 运行镜像地址 **取值范围**： 字符长度为1-128位
	ImageUrl *string `json:"image_url,omitempty"`

	// **参数解释**： SDK检测状态 **取值范围**： - open：开启 - close：关闭
	SdkStatus *string `json:"sdk_status,omitempty"`

	// **参数解释**： 会话开启时间，毫秒时间戳 **取值范围**： 0~9223372036854775807
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 会话上次对话时间，毫秒时间戳 **取值范围**： 0~9223372036854775807
	LastChatTime *int64 `json:"last_chat_time,omitempty"`

	// **参数解释**： 会话运行时长，单位为秒 **取值范围**： 0~2147483647
	RunDuration *int32 `json:"run_duration,omitempty"`

	// **参数解释**： 消耗的token数量 **取值范围**： 0~2147483647
	SecurityTokens *int32 `json:"security_tokens,omitempty"`

	// **参数解释**： 智能体的来源 **取值范围**： - AgentArts - OfficeClaw
	Source *string `json:"source,omitempty"`
}

func (o ListAiSessionsResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiSessionsResponseBodyData struct{}"
	}

	return strings.Join([]string{"ListAiSessionsResponseBodyData", string(data)}, " ")
}
