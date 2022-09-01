package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteApiAclBindingV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 绑定关系编号
	AclBindingsId string `json:"acl_bindings_id" xml:"acl_bindings_id"`
}

func (o DeleteApiAclBindingV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiAclBindingV2Request struct{}"
	}

	return strings.Join([]string{"DeleteApiAclBindingV2Request", string(data)}, " ")
}
