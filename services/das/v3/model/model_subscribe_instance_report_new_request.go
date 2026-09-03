package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeInstanceReportNewRequest Request Object
type SubscribeInstanceReportNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *SubscribeInstanceReportNewRequestBody `json:"body,omitempty"`
}

func (o SubscribeInstanceReportNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeInstanceReportNewRequest struct{}"
	}

	return strings.Join([]string{"SubscribeInstanceReportNewRequest", string(data)}, " ")
}
