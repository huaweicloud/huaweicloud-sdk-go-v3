package model

import (
	"encoding/json"

	"strings"
)

//
type WatermarkTemplate struct {
	// 水印模板名称<br/>

	Name *string `json:"name,omitempty"`
	// 水印模板配置id<br/>

	Id *string `json:"id,omitempty"`
	// 启用状态。  取值为： - 0：停用 - 1：启用

	Status *int64 `json:"status,omitempty"`
	// 水印图片相对输出视频的水平偏移量。  默认值是0。

	Dx *string `json:"dx,omitempty"`
	// 水印图片相对输出视频的垂直偏移量。  默认值是0。

	Dy *string `json:"dy,omitempty"`
	// 水印的位置<br/>

	Position *string `json:"position,omitempty"`
	// 水印图片宽<br/>

	Width *string `json:"width,omitempty"`
	// 水印图片高<br/>

	Height *string `json:"height,omitempty"`
	// 创建时间<br/>

	CreateTime *string `json:"create_time,omitempty"`
	// 水印图片下载url<br/>

	ImageUrl *string `json:"image_url,omitempty"`
	// 水印图片格式类型<br/>

	Type *string `json:"type,omitempty"`
	// 水印类型，当前只支持Image（图片水印）<br/>

	WatermarkType *string `json:"watermark_type,omitempty"`
	// type设置为Image时有效。 目前包括： - Original：只做简单缩放，不做其他处理 - Transparent：图片底色透明 - Grayed：彩色图片变灰

	ImageProcess *string `json:"image_process,omitempty"`
	// 水印开始时间<br/>

	TimelineStart *string `json:"timeline_start,omitempty"`
	// 水印持续时间<br/>

	TimelineDuration *string `json:"timeline_duration,omitempty"`
}

func (o WatermarkTemplate) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "WatermarkTemplate struct{}"
	}

	return strings.Join([]string{"WatermarkTemplate", string(data)}, " ")
}
