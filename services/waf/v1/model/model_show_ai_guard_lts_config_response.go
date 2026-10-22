package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiGuardLtsConfigResponse Response Object
type ShowAiGuardLtsConfigResponse struct {

	// 开关状态（true：开启，false：关闭）
	Enabled *bool `json:"enabled,omitempty"`

	LtsIdInfo      *AiGuardLtsIdInfo `json:"lts_id_info,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowAiGuardLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiGuardLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAiGuardLtsConfigResponse", string(data)}, " ")
}
