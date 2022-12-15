package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResourcesJobDetailRequest struct {

	// 批量操作返回的JOB ID
	JobId string `json:"job_id"`
}

func (o ShowResourcesJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcesJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowResourcesJobDetailRequest", string(data)}, " ")
}
