package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SvgWatermark struct {

	// 水印图片起点相对输出视频顶点的水平偏移量。 设置方法有如下两种： 整数型：表示图片起点水平偏移视频顶点的像素值，单位px。取值范围：[0，4096] 小数型：表示图片起点相对于视频分辨率宽的水平偏移比率。取值范围：(0，1)，支持4位小数，如0.9999，超出部分系统自动丢弃。
	Dx *string `json:"dx,omitempty"`

	// 水印图片起点相对输出视频顶点的垂直偏移量。 设置方法有如下两种： 整数型：表示图片起点水平偏移视频顶点的像素值，单位px。取值范围：[0，4096] 小数型：表示图片起点相对于视频分辨率宽的水平偏移比率。取值范围：(0，1)，支持4位小数，如0.9999，超出部分系统自动丢弃。
	Dy *string `json:"dy,omitempty"`

	// 水印的位置。 取值如下： TopRight：右上角。 TopLeft：左上角。 BottomRight：右下角。 BottomLeft：左下角。 ClockWise：顺时针 AntiClockWise：逆时针 Random：随机跳转
	Referpos *string `json:"referpos,omitempty"`

	// 水印开始时间，与“timeline_duration”配合使用。
	TimelineStart *string `json:"timeline_start,omitempty"`

	// 水印持续时间，与“timeline_start”配合使用。 取值范围：[数字，ToEND]。“ToEND”表示持续到视频结束。 默认值：ToEND。
	TimelineDuration *string `json:"timeline_duration,omitempty"`

	// 随机时间最小值，单位：秒
	RandomTimeMin *float32 `json:"random_time_min,omitempty"`

	// 随机时间最大值，单位：秒
	RandomTimeMax *float32 `json:"random_time_max,omitempty"`

	// 水印图片宽，值有两种形式： 整数型代水印图片宽的像素值，范围[8，4096]，单位px。 小数型代表相对输出视频分辨率宽的比率，范围(0,1)，支持4位小数，如0.9999，超出部分系统自动丢弃。
	Width *string `json:"width,omitempty"`

	// 水印图片高，值有两种形式： 整数型代表水印图片高的像素值，范围[8，4096]，单位px。 小数型代表相对输出视频分辨率高的比率，范围(0，1)，支持4位小数，如0.9999，超出部分系统自动丢弃。
	Height *string `json:"height,omitempty"`
}

func (o SvgWatermark) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SvgWatermark struct{}"
	}

	return strings.Join([]string{"SvgWatermark", string(data)}, " ")
}
