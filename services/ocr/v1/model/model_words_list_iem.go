package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单元格内文字段列表。
type WordsListIem struct {

	// 文字块识别结果。
	Words *string `json:"words,omitempty"`

	// 字段的平均置信度，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *float32 `json:"confidence,omitempty"`

	// 文字块位置信息，列表形式，分别表示文字块4个顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`
}

func (o WordsListIem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WordsListIem struct{}"
	}

	return strings.Join([]string{"WordsListIem", string(data)}, " ")
}
