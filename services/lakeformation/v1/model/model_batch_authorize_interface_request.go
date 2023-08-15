package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAuthorizeInterfaceRequest Request Object
type BatchAuthorizeInterfaceRequest struct {

	// instance id
	InstanceId string `json:"instance_id"`

	Body *AccessPolicyInput `json:"body,omitempty"`
}

func (o BatchAuthorizeInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAuthorizeInterfaceRequest struct{}"
	}

	return strings.Join([]string{"BatchAuthorizeInterfaceRequest", string(data)}, " ")
}
