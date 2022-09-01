package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type FinancialStatementWordsRegionList struct {

	// 区域属性：文本或表格。
	Type *string `json:"type,omitempty" xml:"type"`

	// 区域内文字块数目。对文本区，文字块以文本字段为单位；对表格区，文字块以单元格内所有字段为单位。
	WordsBlockCount *float32 `json:"words_block_count,omitempty" xml:"words_block_count"`

	// 表格位置信息，列表形式，分别表示表格4个顶点的x, y坐标;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TableLocation *[][]int32 `json:"table_location,omitempty" xml:"table_location"`

	// 区域内文字块列表，输出顺序从左到右，从上到下。
	WordsBlockList *[]FinancialStatementWordsBlockList `json:"words_block_list,omitempty" xml:"words_block_list"`
}

func (o FinancialStatementWordsRegionList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FinancialStatementWordsRegionList struct{}"
	}

	return strings.Join([]string{"FinancialStatementWordsRegionList", string(data)}, " ")
}
