package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobRespJob 创建作业的job模型。
type CreateJobRespJob struct {

	// 作业Id。
	JobId *int64 `json:"job_id,omitempty"`

	// 作业状态名称。
	StatusName *string `json:"status_name,omitempty"`

	// 当前状态描述，包含异常状态原因及建议。
	StatusDesc *string `json:"status_desc,omitempty"`
}

func (o CreateJobRespJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobRespJob struct{}"
	}

	return strings.Join([]string{"CreateJobRespJob", string(data)}, " ")
}
