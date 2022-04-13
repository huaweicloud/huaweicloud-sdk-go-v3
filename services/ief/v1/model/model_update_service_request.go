package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateServiceRequest struct {
	// 服务ID

	ServiceId string `json:"service_id"`
	// 铂金版实例ID

	IefInstanceId string `json:"ief-instance-id"`

	Body *Service `json:"body,omitempty"`
}

func (o UpdateServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceRequest struct{}"
	}

	return strings.Join([]string{"UpdateServiceRequest", string(data)}, " ")
}
