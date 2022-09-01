package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SourceInfo struct {

	// 片源时长，单位：秒
	Duration *int32 `json:"duration,omitempty" xml:"duration"`

	// 片源时长，单位：毫秒
	DurationMs *int64 `json:"duration_ms,omitempty" xml:"duration_ms"`

	// 片源格式
	Format *string `json:"format,omitempty" xml:"format"`

	// 片源大小
	Size *int64 `json:"size,omitempty" xml:"size"`

	VideoInfo *VideoInfo `json:"video_info,omitempty" xml:"video_info"`

	// 音频信息
	AudioInfo *[]AudioInfo `json:"audio_info,omitempty" xml:"audio_info"`
}

func (o SourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceInfo struct{}"
	}

	return strings.Join([]string{"SourceInfo", string(data)}, " ")
}
