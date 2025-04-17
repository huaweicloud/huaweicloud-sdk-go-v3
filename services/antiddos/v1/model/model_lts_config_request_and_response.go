package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsConfigRequestAndResponse 日志配置请求
type LtsConfigRequestAndResponse struct {

	// 是否开启日志
	Enabled *bool `json:"enabled,omitempty"`

	LtsIdInfo *LtsConfigRequestAndResponseLtsIdInfo `json:"lts_id_info,omitempty"`
}

func (o LtsConfigRequestAndResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsConfigRequestAndResponse struct{}"
	}

	return strings.Join([]string{"LtsConfigRequestAndResponse", string(data)}, " ")
}
