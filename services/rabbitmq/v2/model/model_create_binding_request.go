package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBindingRequest Request Object
type CreateBindingRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// Vhost名称
	Vhost string `json:"vhost"`

	// Exchange名称
	Exchange string `json:"exchange"`

	Body *CreateBindingBody `json:"body,omitempty"`
}

func (o CreateBindingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBindingRequest struct{}"
	}

	return strings.Join([]string{"CreateBindingRequest", string(data)}, " ")
}
