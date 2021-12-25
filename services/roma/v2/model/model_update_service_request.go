package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateServiceRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 服务ID

	ServiceId string `json:"service_id"`

	Body *UpdateServiceRequestBody `json:"body,omitempty"`
}

func (o UpdateServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceRequest struct{}"
	}

	return strings.Join([]string{"UpdateServiceRequest", string(data)}, " ")
}
