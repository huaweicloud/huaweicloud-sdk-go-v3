package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCancelAuthorizationInterfaceRequest struct {

	// instance id
	InstanceId string `json:"instance_id"`

	Body *AccessPolicyInput `json:"body,omitempty"`
}

func (o BatchCancelAuthorizationInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCancelAuthorizationInterfaceRequest struct{}"
	}

	return strings.Join([]string{"BatchCancelAuthorizationInterfaceRequest", string(data)}, " ")
}
