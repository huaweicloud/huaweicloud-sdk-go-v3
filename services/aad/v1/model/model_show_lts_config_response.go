package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLtsConfigResponse Response Object
type ShowLtsConfigResponse struct {

	// 日志配置开关
	Enabled *bool `json:"enabled,omitempty"`

	LtsIdInfo      *UpdateLtsConfigRequestBodyLtsIdInfo `json:"lts_id_info,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowLtsConfigResponse", string(data)}, " ")
}
