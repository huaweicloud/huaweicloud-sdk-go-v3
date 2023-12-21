package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobDetailRequest Request Object
type ShowJobDetailRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ShowJobDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowJobDetailRequest", string(data)}, " ")
}
