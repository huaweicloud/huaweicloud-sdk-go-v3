package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBatchJobResponse Response Object
type CreateBatchJobResponse struct {

	// 批量处理作业ID
	JobId *string `json:"job_id,omitempty"`

	// 批量处理作业名称
	JobName *string `json:"job_name,omitempty"`

	// 批量处理作业创建时间戳
	CreatedAt      *int32 `json:"created_at,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchJobResponse struct{}"
	}

	return strings.Join([]string{"CreateBatchJobResponse", string(data)}, " ")
}
