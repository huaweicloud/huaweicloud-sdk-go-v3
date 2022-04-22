package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCommandRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 服务ID
	ServiceId string `json:"service_id"`

	Body *CreateCommandRequestBody `json:"body,omitempty"`
}

func (o CreateCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommandRequest struct{}"
	}

	return strings.Join([]string{"CreateCommandRequest", string(data)}, " ")
}
