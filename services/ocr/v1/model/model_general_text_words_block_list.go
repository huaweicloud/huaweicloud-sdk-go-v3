package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 识别文字块列表，输出顺序从左到右，先上后下。
type GeneralTextWordsBlockList struct {

	// 文字块识别结果。
	Words string `json:"words" xml:"words"`

	// 文字块的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location [][]int32 `json:"location" xml:"location"`
}

func (o GeneralTextWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralTextWordsBlockList struct{}"
	}

	return strings.Join([]string{"GeneralTextWordsBlockList", string(data)}, " ")
}
