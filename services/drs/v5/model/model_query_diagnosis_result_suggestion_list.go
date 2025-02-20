package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDiagnosisResultSuggestionList struct {

	// 诊断建议内容。
	Content *string `json:"content,omitempty"`
}

func (o QueryDiagnosisResultSuggestionList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDiagnosisResultSuggestionList struct{}"
	}

	return strings.Join([]string{"QueryDiagnosisResultSuggestionList", string(data)}, " ")
}
