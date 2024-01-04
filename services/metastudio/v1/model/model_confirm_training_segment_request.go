package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmTrainingSegmentRequest Request Object
type ConfirmTrainingSegmentRequest struct {

	// 任务id。
	JobId string `json:"job_id"`

	// 语句索引。
	Index int32 `json:"index"`
}

func (o ConfirmTrainingSegmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmTrainingSegmentRequest struct{}"
	}

	return strings.Join([]string{"ConfirmTrainingSegmentRequest", string(data)}, " ")
}
