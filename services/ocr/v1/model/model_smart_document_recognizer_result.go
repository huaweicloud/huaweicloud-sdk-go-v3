package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerResult struct {
	OcrResult *SmartDocumentRecognizerOcrResult `json:"ocr_result"`

	KvResult *SmartDocumentRecognizerKvResult `json:"kv_result,omitempty"`

	TableResult *SmartDocumentRecognizerTableResult `json:"table_result,omitempty"`

	LayoutResult *SmartDocumentRecognizerLayoutResult `json:"layout_result,omitempty"`

	FormResult *SmartDocumentRecognizerFormResult `json:"form_result,omitempty"`

	FormulaResult *SmartDocumentRecognizerFormulaResult `json:"formula_result,omitempty"`
}

func (o SmartDocumentRecognizerResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerResult struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerResult", string(data)}, " ")
}
