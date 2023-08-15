package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRunFlinkJobsRequestBody
type BatchRunFlinkJobsRequestBody struct {

	// 是否将作业从最近创建的保存点恢复。类型为boolean。  当resume_savepoint为true时，表示作业从最近创建的保存点恢复。 当resume_savepoint为false时，表示不恢复正常启动。默认为false。
	ResumeSavepoint *bool `json:"resume_savepoint,omitempty"`

	//
	JobIds []int64 `json:"job_ids"`
}

func (o BatchRunFlinkJobsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRunFlinkJobsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchRunFlinkJobsRequestBody", string(data)}, " ")
}
