package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OriginPara struct {

	// 片源时长，单位：秒
	Duration *int32 `json:"duration,omitempty" xml:"duration"`

	// 片源时长，单位：毫秒
	DurationMs *int64 `json:"duration_ms,omitempty" xml:"duration_ms"`

	// 文件格式
	FileFormat *string `json:"file_format,omitempty" xml:"file_format"`

	Video *VideoInfo `json:"video,omitempty" xml:"video"`

	Audio *AudioInfo `json:"audio,omitempty" xml:"audio"`
}

func (o OriginPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginPara struct{}"
	}

	return strings.Join([]string{"OriginPara", string(data)}, " ")
}
