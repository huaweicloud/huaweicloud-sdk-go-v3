package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaseCompletionRateVo 计算用例完成率
type CaseCompletionRateVo struct {

	// 总用例数
	Total *int32 `json:"total,omitempty"`

	// 用例完成率
	CompletionRate *string `json:"completion_rate,omitempty"`

	// 用户自定义状态对应的用例数目
	StatusNumberList *[]NameAndValueVo `json:"status_number_list,omitempty"`
}

func (o CaseCompletionRateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseCompletionRateVo struct{}"
	}

	return strings.Join([]string{"CaseCompletionRateVo", string(data)}, " ")
}
