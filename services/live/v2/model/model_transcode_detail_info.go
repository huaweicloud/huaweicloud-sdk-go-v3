package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TranscodeDetailInfo struct {

	// 流名
	StreamName *string `json:"stream_name,omitempty"`

	// 转码模板
	Template *string `json:"template,omitempty"`

	// 转码格式（H264/H265）
	CodeRateFormat *string `json:"code_rate_format,omitempty"`

	// 转码时长
	Duration *int64 `json:"duration,omitempty"`
}

func (o TranscodeDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscodeDetailInfo struct{}"
	}

	return strings.Join([]string{"TranscodeDetailInfo", string(data)}, " ")
}
