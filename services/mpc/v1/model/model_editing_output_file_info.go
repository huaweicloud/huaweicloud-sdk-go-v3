package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditingOutputFileInfo struct {

	// 输出片源时长 单位秒
	Duration *int64 `json:"duration,omitempty"`

	// 输出封装格式
	Format *string `json:"format,omitempty"`

	// 媒体文件大小，单位：KByte。
	Size *int64 `json:"size,omitempty"`

	// 媒体文件大小，单位：Byte。
	SizeByte *int64 `json:"size_byte,omitempty"`

	VideoInfo *EditingVideoInfo `json:"video_info,omitempty"`

	// 音频信息
	AudioInfo *[]EditingAudioInfo `json:"audio_info,omitempty"`
}

func (o EditingOutputFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditingOutputFileInfo struct{}"
	}

	return strings.Join([]string{"EditingOutputFileInfo", string(data)}, " ")
}
