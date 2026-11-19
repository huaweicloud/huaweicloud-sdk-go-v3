package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyDetailResponseInfo 策略详情
type ShowPolicyDetailResponseInfo struct {

	// **参数解释**： 功能类别 **取值范围**： 字符长度1-128位
	Category *string `json:"category,omitempty"`

	// **参数解释**： 策略内容 **取值范围**： 字符长度0-4096位
	Content *string `json:"content,omitempty"`

	// **参数解释**： 策略描述 **取值范围**： 字符长度0-1024位
	Description *string `json:"description,omitempty"`

	// **参数解释**： 策略是否开启 **取值范围**： - true：开启 - false：关闭
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**： 策略名 **取值范围**： 字符长度1-256位
	Name *string `json:"name,omitempty"`

	// **参数解释**： 策略支持的主机种类 **取值范围**: - Linux：Linux操作系统 - Windows：Windows操作系统
	SupportOs *string `json:"support_os,omitempty"`

	// **参数解释**： 策略Id **取值范围**： 字符长度1-128位
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**： 特性名 **取值范围**： 字符长度1-128位
	FeatureName *string `json:"feature_name,omitempty"`
}

func (o ShowPolicyDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"ShowPolicyDetailResponseInfo", string(data)}, " ")
}
