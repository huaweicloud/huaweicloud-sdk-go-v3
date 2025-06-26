package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisSummaryRequest Request Object
type ShowDiagnosisSummaryRequest struct {

	// 诊断任务ID
	TaskId string `json:"task_id"`
}

func (o ShowDiagnosisSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisSummaryRequest", string(data)}, " ")
}
