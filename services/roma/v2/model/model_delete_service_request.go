package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceRequest Request Object
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
