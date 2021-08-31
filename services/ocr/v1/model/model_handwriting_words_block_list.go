package model

import (
	"encoding/json"

	"strings"
)

//
type HandwritingWordsBlockList struct {
	// 文字块识别结果。

	Words *string `json:"words,omitempty"`
	// 说明该识别结果所属类型，例如：handwriting。

	Type *string `json:"type,omitempty"`
	// 文字块words的置信度。

	Confidence *float32 `json:"confidence,omitempty"`
	// 文字块words的区域位置信息，列表形式，分别表示文字块顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。

	Location *[][]int32 `json:"location,omitempty"`
}

func (o HandwritingWordsBlockList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HandwritingWordsBlockList struct{}"
	}

	return strings.Join([]string{"HandwritingWordsBlockList", string(data)}, " ")
}
