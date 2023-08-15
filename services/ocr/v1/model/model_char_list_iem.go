package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CharListIem struct {

	// 单字符识别结果。
	Char *string `json:"char,omitempty"`

	// 单字符的置信度，置信度越大，表示本次识别的文字的可靠性越高，在统计意义上，置信度越大，准确率越高。置信度由算法给出，不直接等价于对应字段的准确率。
	CharConfidence *float32 `json:"char_confidence,omitempty"`

	// 单字符的位置信息，列表形式，分别表示文字块4个顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	CharLocation *[][]int32 `json:"char_location,omitempty"`
}

func (o CharListIem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CharListIem struct{}"
	}

	return strings.Join([]string{"CharListIem", string(data)}, " ")
}
