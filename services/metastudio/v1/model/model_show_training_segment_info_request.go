package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingSegmentInfoRequest Request Object
type ShowTrainingSegmentInfoRequest struct {

	// 任务id。
	JobId string `json:"job_id"`
}

func (o ShowTrainingSegmentInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingSegmentInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainingSegmentInfoRequest", string(data)}, " ")
}
