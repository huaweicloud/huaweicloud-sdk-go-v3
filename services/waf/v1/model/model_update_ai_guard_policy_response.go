package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiGuardPolicyResponse Response Object
type UpdateAiGuardPolicyResponse struct {

	// 策略 ID
	Id *string `json:"id,omitempty"`

	// 策略名称
	Name *string `json:"name,omitempty"`

	// 绑定的 AI 应用列表
	BindAiApps *[]AiAppBasicInfo `json:"bind_ai_apps,omitempty"`

	// 创建时间
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateAiGuardPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardPolicyResponse", string(data)}, " ")
}
