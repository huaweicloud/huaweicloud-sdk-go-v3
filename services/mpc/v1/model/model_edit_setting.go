package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditSetting struct {

	// input指定源文件开始截取的时间，单位：秒。可以有正数或者负数，正数表示从开始往后的时间点，负数表示从结尾往前的时间点
	TimelineStart *string `json:"timeline_start,omitempty" xml:"timeline_start"`

	// input指定源文件接受截取的时间，单位：秒。可以有正数或者负数，正数表示从开始往后的时间点，负数表示从结尾往前的时间点。
	TimelineEnd *string `json:"timeline_end,omitempty" xml:"timeline_end"`

	// 转码模板id
	TransTemplateId *int32 `json:"trans_template_id,omitempty" xml:"trans_template_id"`

	AvParameter *AvParameters `json:"av_parameter,omitempty" xml:"av_parameter"`

	// 马赛克（模糊处理）配置，会对input指定的源文件进行马赛克处理，马赛克基于视频左上角为参考位置
	Mosaics *[]MosaicInfo `json:"mosaics,omitempty" xml:"mosaics"`

	// 图片水印配置，会对input指定的源文件进行马赛克处理。水印设置参数里面的overlay_input字段不填
	ImageWatermarks *[]ImageWatermarkSetting `json:"image_watermarks,omitempty" xml:"image_watermarks"`

	// 头部文件列表，需要指定文件名。列表内文件会按照顺序拼接在input指定文件的前面
	Heads *[]ObsObjInfo `json:"heads,omitempty" xml:"heads"`

	// 尾部文件列表，需要指定文件名。列表内文件会按照顺序拼接在input指定文件的后面
	Tails *[]ObsObjInfo `json:"tails,omitempty" xml:"tails"`

	Output *ObsObjInfo `json:"output" xml:"output"`
}

func (o EditSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditSetting struct{}"
	}

	return strings.Join([]string{"EditSetting", string(data)}, " ")
}
