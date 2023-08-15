package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePropertyRequest Request Object
type CreatePropertyRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id"`

	Body *CreatePropertyRequestBody `json:"body,omitempty"`
}

func (o CreatePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePropertyRequest struct{}"
	}

	return strings.Join([]string{"CreatePropertyRequest", string(data)}, " ")
}
