package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartDocumentRecognizerFormResult struct {

	// 模型识别到的有线表单数量。
	FormCount *int32 `json:"form_count,omitempty"`

	// 有线表单识别结果列表。
	FormList *[]SmartDocumentRecognizerTableBlock `json:"form_list,omitempty"`
}

func (o SmartDocumentRecognizerFormResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartDocumentRecognizerFormResult struct{}"
	}

	return strings.Join([]string{"SmartDocumentRecognizerFormResult", string(data)}, " ")
}
