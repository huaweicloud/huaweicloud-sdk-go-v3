package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclBindingBatchDelete struct {

	// 需要解除绑定的API和ACL绑定关系ID列表
	AclBindings *[]string `json:"acl_bindings,omitempty" xml:"acl_bindings"`
}

func (o AclBindingBatchDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclBindingBatchDelete struct{}"
	}

	return strings.Join([]string{"AclBindingBatchDelete", string(data)}, " ")
}
