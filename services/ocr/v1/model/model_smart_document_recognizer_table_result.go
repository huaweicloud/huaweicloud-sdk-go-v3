package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerTableResult struct {

	// 模型识别到的表格数量。
	TableCount *int32 `json:"table_count,omitempty"`

	// 表格识别结果列表。
	TableList *[]SmartDocumentRecognizerTableBlock `json:"table_list,omitempty"`
}

func (o SmartDocumentRecognizerTableResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerTableResult struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerTableResult", string(data)}, " ")
}
