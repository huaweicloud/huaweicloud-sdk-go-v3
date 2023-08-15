package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBatchJobResponse Response Object
type ListBatchJobResponse struct {

	// 数目
	JobCount *int32 `json:"job_count,omitempty"`

	// 批量处理作业详情
	Jobs           *[]BatchJobForList `json:"jobs,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListBatchJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBatchJobResponse struct{}"
	}

	return strings.Join([]string{"ListBatchJobResponse", string(data)}, " ")
}
