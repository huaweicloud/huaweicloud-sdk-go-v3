package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerFormulaResult struct {

	// 数学公式数量。
	FormulaCount *int32 `json:"formula_count,omitempty"`

	// 数学公式识别结果列表。
	FormulaList *[]SmartDocumentRecognizerFormulaBlock `json:"formula_list,omitempty"`
}

func (o SmartDocumentRecognizerFormulaResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerFormulaResult struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerFormulaResult", string(data)}, " ")
}
