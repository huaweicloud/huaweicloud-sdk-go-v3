package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateServiceRequest struct {

	// 铂金版实例ID
	IefInstanceId string `json:"ief-instance-id"`

	Body *Service `json:"body,omitempty"`
}

func (o CreateServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceRequest struct{}"
	}

	return strings.Join([]string{"CreateServiceRequest", string(data)}, " ")
}
