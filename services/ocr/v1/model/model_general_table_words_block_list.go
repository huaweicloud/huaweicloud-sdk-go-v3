package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 子区域识别文字块列表，输出顺序从左到右，先上后下。
type GeneralTableWordsBlockList struct {
	// 文字块识别结果。

	Words *string `json:"words,omitempty"`
	// 字段的平均置信度，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。置信度由算法给出，不直接等价于对应字段的准确率。

	Confidence *float32 `json:"confidence,omitempty"`
	// 文字块位置信息，列表形式，分别表示文字块4个顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。

	Location *[][]int32 `json:"location,omitempty"`
	// 单元格内文字段列表。输出顺序从左到右，从上到下。仅当入参\"return_text_location\"为true时存在。

	WordsList *[]WordsListIem `json:"words_list,omitempty"`
	// 文字块占用的行信息，编号从0开始，列表形式，数据类型为Integer。仅在表格区域内有效，即type字段为\"table\"时该字段有效。

	Rows *[]int32 `json:"rows,omitempty"`
	// 文字块占用的列信息，编号从0开始，列表形式，数据类型为Integer。仅在表格区域内有效，即type字段为\"table\"时该字段有效。

	Columns *[]int32 `json:"columns,omitempty"`
	// 单元格位置信息，列表形式，分别表示单元格4个顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。

	CellLocation *[][]int32 `json:"cell_location,omitempty"`
}

func (o GeneralTableWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralTableWordsBlockList struct{}"
	}

	return strings.Join([]string{"GeneralTableWordsBlockList", string(data)}, " ")
}
