package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclAccountRoleRequest Request Object
type UpdateAclAccountRoleRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// ACL账号ID。
	AccountId string `json:"account_id"`

	Body *AclAccountRoleModifyBody `json:"body,omitempty"`
}

func (o UpdateAclAccountRoleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclAccountRoleRequest struct{}"
	}

	return strings.Join([]string{"UpdateAclAccountRoleRequest", string(data)}, " ")
}
