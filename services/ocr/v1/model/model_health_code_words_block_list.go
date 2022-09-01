package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type HealthCodeWordsBlockList struct {

	// 文字块识别结果。
	Words *string `json:"words,omitempty" xml:"words"`

	// 识别到的文字块的区域位置信息，列表形式，分别表示文字块4个顶点的（x,y）坐标。采用图像坐标系，图像坐标原点为图像左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty" xml:"location"`

	// 各个字段的置信度
	Confidence *float32 `json:"confidence,omitempty" xml:"confidence"`
}

func (o HealthCodeWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCodeWordsBlockList struct{}"
	}

	return strings.Join([]string{"HealthCodeWordsBlockList", string(data)}, " ")
}
