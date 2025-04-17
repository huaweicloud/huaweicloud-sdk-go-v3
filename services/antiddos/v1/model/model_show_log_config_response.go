package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogConfigResponse Response Object
type ShowLogConfigResponse struct {

	// 是否开启日志
	Enabled *bool `json:"enabled,omitempty"`

	LtsIdInfo      *LtsConfigRequestAndResponseLtsIdInfo `json:"lts_id_info,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ShowLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowLogConfigResponse", string(data)}, " ")
}
