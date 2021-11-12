package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AnswerBlockList struct {
	// 文字块识别结果。

	Words string `json:"words"`
	// 文字块words的置信度。

	Confidence float32 `json:"confidence"`
	// 文字块words的区域位置信息，列表形式，分别表示文字块顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。

	Location []int32 `json:"location"`
}

func (o AnswerBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnswerBlockList struct{}"
	}

	return strings.Join([]string{"AnswerBlockList", string(data)}, " ")
}
