package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListJobDetailsRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o ListJobDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListJobDetailsRequest", string(data)}, " ")
}
