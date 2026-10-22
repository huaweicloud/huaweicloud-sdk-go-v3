package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateBotMTenantPolicyRequestBody struct {
	UserTag *UserTagDto `json:"user_tag,omitempty"`

	// 场景标注（登录/注册场景接口的语义标注列表）
	ScenarioAnnotations *[]ScenarioAnnotationDto `json:"scenario_annotations,omitempty"`
}

func (o UpdateBotMTenantPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMTenantPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBotMTenantPolicyRequestBody", string(data)}, " ")
}
