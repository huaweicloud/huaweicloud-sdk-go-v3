package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobProgressRequest struct {

	// 异步任务ID
	JobId string `json:"job_id"`
}

func (o ShowJobProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowJobProgressRequest", string(data)}, " ")
}
