package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type FinancialStatementWordsBlockList struct {
	// 文字块内容。当入参\"return_text_location\"为false时，每个单元格返回一个文本值，不同行文本由换行符 \"\\n\" 拼接。

	Words *string `json:"words,omitempty"`
	// 文字块位置信息，列表形式，分别表示文字块4个顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。

	Location *[][]int32 `json:"location,omitempty"`
	// 文字块识别结果置信度信息，置信度越大，表示本次识别的对应字段的可靠性越大，在统计意义上，置信度越大正确率越高。注：置信度由算法给出，其不直接等价于对应字段的精度。

	Confidence *float32 `json:"confidence,omitempty"`
	// 单元格行信息，列表形式。多个连续值表示单元格垮多行。

	Rows *[]int32 `json:"rows,omitempty"`
	// 单元格列信息，列表形式。多个连续值表示单元格垮多列。

	Columns *[]int32 `json:"columns,omitempty"`
	// 单元格位置信息，列表形式，分别表示单元格4个顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。

	CellLocation *[][]int32 `json:"cell_location,omitempty"`
}

func (o FinancialStatementWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FinancialStatementWordsBlockList struct{}"
	}

	return strings.Join([]string{"FinancialStatementWordsBlockList", string(data)}, " ")
}
