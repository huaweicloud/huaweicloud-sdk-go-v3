package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePolicyRequestBody struct {

	// 策略名称（策略名称只能由数字、字母和下划线组成，长度不能超过64为字符）
	Name string `json:"name"`

	// cc规则和精准防护规则“防护动作”选择“仅记录”时，Web基础防护是否命中策略规则并阻断，默认为true
	LogActionReplaced *bool `json:"log_action_replaced,omitempty"`
}

func (o CreatePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePolicyRequestBody", string(data)}, " ")
}
