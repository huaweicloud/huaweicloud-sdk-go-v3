package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GeneralTextCharList struct {

	// 单字符识别结果。
	Char *string `json:"char,omitempty"`

	// 单字符的区域位置信息，列表形式，包含字符区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	CharLocation *[][]int32 `json:"char_location,omitempty"`

	// 单字符识别结果的置信度。
	CharConfidence *float32 `json:"char_confidence,omitempty"`
}

func (o GeneralTextCharList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralTextCharList struct{}"
	}

	return strings.Join([]string{"GeneralTextCharList", string(data)}, " ")
}
