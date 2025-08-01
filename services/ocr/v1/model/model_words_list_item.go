package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WordsListItem 单元格内文字段列表。
type WordsListItem struct {

	// 文字块识别结果。
	Words *string `json:"words,omitempty"`

	// 字段的平均置信度，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *float32 `json:"confidence,omitempty"`

	// 文字块位置信息，列表形式，分别表示文字块4个顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`

	// 单元格内文字段列表。输出顺序从左到右，从上到下。仅当入参\"return_text_location\"和\"return_char_location\"同时为true时存在。
	CharList *[]CharListItem `json:"char_list,omitempty"`
}

func (o WordsListItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WordsListItem struct{}"
	}

	return strings.Join([]string{"WordsListItem", string(data)}, " ")
}
