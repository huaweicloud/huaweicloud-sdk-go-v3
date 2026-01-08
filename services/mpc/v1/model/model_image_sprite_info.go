package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageSpriteInfo struct {

	// 雪碧图中小图的行数。
	RowCount *int32 `json:"row_count,omitempty"`

	// 雪碧图中小图的列数。
	ColumnCount *int32 `json:"column_count,omitempty"`

	// 雪碧图中小图数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 雪碧图小图宽度
	Width *int32 `json:"width,omitempty"`

	// 雪碧图小图高度
	Height *int32 `json:"height,omitempty"`

	// 每一张雪碧图大图的路径。
	OutputImageName *[]string `json:"output_image_name,omitempty"`

	// 雪碧图子图位置与时间关系的 WebVtt 文件路径。WebVtt 文件表明了各个雪碧图小图对应的时间点，以及在雪碧大图里的坐标位置，一般被播放器用于实现预览。
	OutputWebvttName *string `json:"output_webvtt_name,omitempty"`

	Output *ObsObjInfo `json:"output,omitempty"`
}

func (o ImageSpriteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageSpriteInfo struct{}"
	}

	return strings.Join([]string{"ImageSpriteInfo", string(data)}, " ")
}
