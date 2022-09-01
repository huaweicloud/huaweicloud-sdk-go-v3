package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBatchJobRequest struct {

	// 数据开发任务ID。
	JobId string `json:"job_id" xml:"job_id"`
}

func (o ShowBatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchJobRequest struct{}"
	}

	return strings.Join([]string{"ShowBatchJobRequest", string(data)}, " ")
}
