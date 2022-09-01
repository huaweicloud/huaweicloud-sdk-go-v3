package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OutputVideoPara struct {

	// 输出视频对应的模板ID
	TemplateId *int32 `json:"template_id,omitempty" xml:"template_id"`

	// 视频大小
	Size *int64 `json:"size,omitempty" xml:"size"`

	// 视频封装格式
	Pack *string `json:"pack,omitempty" xml:"pack"`

	Video *VideoInfo `json:"video,omitempty" xml:"video"`

	Audio *AudioInfo `json:"audio,omitempty" xml:"audio"`

	// 输出片源文件名
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 折算后视频时长
	ConverDuration *float64 `json:"conver_duration,omitempty" xml:"conver_duration"`

	Error *XCodeError `json:"error,omitempty" xml:"error"`
}

func (o OutputVideoPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutputVideoPara struct{}"
	}

	return strings.Join([]string{"OutputVideoPara", string(data)}, " ")
}
