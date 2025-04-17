package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LtsConfigRequestAndResponseLtsIdInfo 日志信息
type LtsConfigRequestAndResponseLtsIdInfo struct {

	// 日志组id
	LtsGroupId *string `json:"lts_group_id,omitempty"`

	// 日志流id
	LtsAttackStreamId *string `json:"lts_attack_stream_id,omitempty"`
}

func (o LtsConfigRequestAndResponseLtsIdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsConfigRequestAndResponseLtsIdInfo struct{}"
	}

	return strings.Join([]string{"LtsConfigRequestAndResponseLtsIdInfo", string(data)}, " ")
}
