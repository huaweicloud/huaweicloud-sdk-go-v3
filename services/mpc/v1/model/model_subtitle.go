package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Subtitle struct {
	Input *ObsObjInfo `json:"input,omitempty" xml:"input"`

	// 多字幕文件地址。
	Inputs *[]MulInputFileInfo `json:"inputs,omitempty" xml:"inputs"`

	// 字幕类型。取值如下：  - 0，表示不输出字幕 - 1，表示外部字幕文件嵌入视频流 - 2，表示输出WebVTT格式字幕
	SubtitleType *int32 `json:"subtitle_type,omitempty" xml:"subtitle_type"`
}

func (o Subtitle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subtitle struct{}"
	}

	return strings.Join([]string{"Subtitle", string(data)}, " ")
}
