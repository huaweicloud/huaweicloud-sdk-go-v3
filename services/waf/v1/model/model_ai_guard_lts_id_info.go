package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiGuardLtsIdInfo LTS 日志组和流 ID 信息
type AiGuardLtsIdInfo struct {

	// LTS 日志流 ID
	LtsAccessStreamId *string `json:"lts_access_stream_id,omitempty"`

	// LTS 日志组 ID
	LtsGroupId *string `json:"lts_group_id,omitempty"`
}

func (o AiGuardLtsIdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiGuardLtsIdInfo struct{}"
	}

	return strings.Join([]string{"AiGuardLtsIdInfo", string(data)}, " ")
}
