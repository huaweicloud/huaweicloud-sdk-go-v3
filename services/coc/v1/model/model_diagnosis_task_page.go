package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiagnosisTaskPage 诊断记录分页对象
type DiagnosisTaskPage struct {

	// 一共有多少条数据
	Total *int32 `json:"total,omitempty"`

	// data
	Data *[]DiagnosisTask `json:"data,omitempty"`
}

func (o DiagnosisTaskPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisTaskPage struct{}"
	}

	return strings.Join([]string{"DiagnosisTaskPage", string(data)}, " ")
}
