package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteBatchJobRequest struct {

	// 数据开发任务ID。
	JobId string `json:"job_id"`
}

func (o DeleteBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBatchJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteBatchJobRequest", string(data)}, " ")
}
