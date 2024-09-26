package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHttpPolicyRequestBody struct {

	// 防护策略名
	Name *string `json:"name,omitempty"`

	Action *HttpPolicyAction `json:"action,omitempty"`

	Options *HttpPolicyOption `json:"options,omitempty"`

	// 防护等级
	Level *int32 `json:"level,omitempty"`

	// 精准防护中的检测模式
	FullDetection *bool `json:"full_detection,omitempty"`

	RobotAction *HttpPolicyAction `json:"robot_action,omitempty"`

	ThirdBotOptions *HttpThirdBotOptions `json:"third_bot_options,omitempty"`

	// 扩展字段
	Extend map[string]string `json:"extend,omitempty"`
}

func (o UpdateHttpPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpPolicyRequestBody", string(data)}, " ")
}
