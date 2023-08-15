package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopFlinkJobsRequestBody
type StopFlinkJobsRequestBody struct {

	// 在停止作业之前，用户可以选择是否对作业创建保存点，保存作业的状态信息。类型为boolean。  当triggerSavePoint为true时，表示创建保存点。 当triggerSavePoint为false时，表示不创建保存点。默认为false。
	TriggerSavepoint *bool `json:"trigger_savepoint,omitempty"`

	//
	JobIds []int64 `json:"job_ids"`
}

func (o StopFlinkJobsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopFlinkJobsRequestBody struct{}"
	}

	return strings.Join([]string{"StopFlinkJobsRequestBody", string(data)}, " ")
}
