package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteServiceRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 服务ID

	ServiceId string `json:"service_id"`
}

func (o DeleteServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceRequest struct{}"
	}

	return strings.Join([]string{"DeleteServiceRequest", string(data)}, " ")
}
