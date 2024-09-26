package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowHttpPolicyResponseBody struct {

	// 防护策略id
	Id *string `json:"id,omitempty"`

	// 防护策略名
	Name *string `json:"name,omitempty"`

	// 防护等级
	Level *int32 `json:"level,omitempty"`

	// 精准防护中的检测模式
	FullDetection *bool `json:"full_detection,omitempty"`

	Action *HttpPolicyAction `json:"action,omitempty"`

	RobotAction *HttpPolicyAction `json:"robot_action,omitempty"`

	Options *HttpPolicyOption `json:"options,omitempty"`

	// 防护域名的信息
	BindHost *[]HttpPolicyBindHost `json:"bind_host,omitempty"`

	// 扩展字段
	Extend map[string]string `json:"extend,omitempty"`

	// 三方BOT操作
	ThirdBotOptions map[string]interface{} `json:"third_bot_options,omitempty"`

	// web基础防护托管规则集id
	WapManagedRulesetId *string `json:"wap_managed_ruleset_id,omitempty"`
}

func (o ShowHttpPolicyResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpPolicyResponseBody struct{}"
	}

	return strings.Join([]string{"ShowHttpPolicyResponseBody", string(data)}, " ")
}
