package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeInstanceReportNewRequest Request Object
type UnsubscribeInstanceReportNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *UnsubscribeInstanceReportNewRequestBody `json:"body,omitempty"`
}

func (o UnsubscribeInstanceReportNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeInstanceReportNewRequest struct{}"
	}

	return strings.Join([]string{"UnsubscribeInstanceReportNewRequest", string(data)}, " ")
}
