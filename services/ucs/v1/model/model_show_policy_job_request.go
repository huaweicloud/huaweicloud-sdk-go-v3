package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyJobRequest Request Object
type ShowPolicyJobRequest struct {

	// 任务id
	Jobid string `json:"jobid"`
}

func (o ShowPolicyJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyJobRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyJobRequest", string(data)}, " ")
}
