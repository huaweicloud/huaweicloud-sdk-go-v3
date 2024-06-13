package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaseExecuteVo 需求关联用例执行率情况
type CaseExecuteVo struct {

	// 需求关联用例执行率
	ExecuteRate *string `json:"execute_rate,omitempty"`

	// 需求关联已执行用例总数
	ExecutedNumber *int32 `json:"executed_number,omitempty"`

	// 需求关联未执行用例总数
	NotExecutedNumber *int32 `json:"not_executed_number,omitempty"`
}

func (o CaseExecuteVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseExecuteVo struct{}"
	}

	return strings.Join([]string{"CaseExecuteVo", string(data)}, " ")
}
