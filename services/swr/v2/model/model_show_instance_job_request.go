package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceJobRequest Request Object
type ShowInstanceJobRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o ShowInstanceJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceJobRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceJobRequest", string(data)}, " ")
}
