package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiGuardLtsConfigRequestBody LTS 配置请求/响应体
type UpdateAiGuardLtsConfigRequestBody struct {

	// 开关状态（true：开启，false：关闭）
	Enabled *bool `json:"enabled,omitempty"`

	LtsIdInfo *AiGuardLtsIdInfo `json:"lts_id_info,omitempty"`
}

func (o UpdateAiGuardLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardLtsConfigRequestBody", string(data)}, " ")
}
