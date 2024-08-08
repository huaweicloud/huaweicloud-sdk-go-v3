package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageJobRequest Request Object
type ShowImageJobRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ShowImageJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageJobRequest struct{}"
	}

	return strings.Join([]string{"ShowImageJobRequest", string(data)}, " ")
}
