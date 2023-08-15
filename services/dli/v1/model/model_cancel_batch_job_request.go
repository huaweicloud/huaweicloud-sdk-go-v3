package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelBatchJobRequest Request Object
type CancelBatchJobRequest struct {

	// 批处理作业的ID。
	BatchId string `json:"batch_id"`
}

func (o CancelBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelBatchJobRequest struct{}"
	}

	return strings.Join([]string{"CancelBatchJobRequest", string(data)}, " ")
}
