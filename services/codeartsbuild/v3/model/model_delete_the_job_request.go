package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTheJobRequest Request Object
type DeleteTheJobRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o DeleteTheJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTheJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteTheJobRequest", string(data)}, " ")
}
