package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApiAclBindingV2Request Request Object
type CreateApiAclBindingV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *AclApiBindingCreate `json:"body,omitempty"`
}

func (o CreateApiAclBindingV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiAclBindingV2Request struct{}"
	}

	return strings.Join([]string{"CreateApiAclBindingV2Request", string(data)}, " ")
}
