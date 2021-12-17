package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateBatchJobRequest struct {
	Body *Job `json:"body,omitempty"`
}

func (o CreateBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchJobRequest struct{}"
	}

	return strings.Join([]string{"CreateBatchJobRequest", string(data)}, " ")
}
