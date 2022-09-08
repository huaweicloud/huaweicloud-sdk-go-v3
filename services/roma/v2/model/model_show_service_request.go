package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowServiceRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id"`
}

func (o ShowServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServiceRequest struct{}"
	}

	return strings.Join([]string{"ShowServiceRequest", string(data)}, " ")
}
