package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiGuardPolicyItem 策略详情
type AiGuardPolicyItem struct {

	// 策略 ID
	Id *string `json:"id,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 绑定的 AI 应用列表
	BindAiApps *[]AiAppBasicInfo `json:"bind_ai_apps,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o AiGuardPolicyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiGuardPolicyItem struct{}"
	}

	return strings.Join([]string{"AiGuardPolicyItem", string(data)}, " ")
}
