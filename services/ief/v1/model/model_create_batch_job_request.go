package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateBatchJobRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *BatchJobRequest `json:"body,omitempty"`
}

func (o CreateBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchJobRequest struct{}"
	}

	return strings.Join([]string{"CreateBatchJobRequest", string(data)}, " ")
}
