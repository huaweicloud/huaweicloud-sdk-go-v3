package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiagnoseRecordsRequest Request Object
type ListDiagnoseRecordsRequest struct {

	// 单页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 任务类型，例如 ecs诊断任务 1，rds诊断任务 2
	Type *int32 `json:"type,omitempty"`
}

func (o ListDiagnoseRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnoseRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListDiagnoseRecordsRequest", string(data)}, " ")
}
