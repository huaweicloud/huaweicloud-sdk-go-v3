package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryDiagnosisResultDiagnosisResults struct {
	Item *QueryDiagnosisResultItem `json:"item,omitempty"`

	// 诊断项状态。
	State *string `json:"state,omitempty"`

	// 诊断结果。
	ResultList *[]QueryDiagnosisResultResultList `json:"result_list,omitempty"`

	// 诊断建议。
	SuggestionList *[]QueryDiagnosisResultSuggestionList `json:"suggestion_list,omitempty"`

	// 诊断项得分。
	Score *int32 `json:"score,omitempty"`

	// 完成时间。
	Time *string `json:"time,omitempty"`
}

func (o QueryDiagnosisResultDiagnosisResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDiagnosisResultDiagnosisResults struct{}"
	}

	return strings.Join([]string{"QueryDiagnosisResultDiagnosisResults", string(data)}, " ")
}
