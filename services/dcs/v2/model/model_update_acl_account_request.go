package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclAccountRequest Request Object
type UpdateAclAccountRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// ACL账号ID。
	AccountId string `json:"account_id"`

	Body *AclAccountRoleModifyBody `json:"body,omitempty"`
}

func (o UpdateAclAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclAccountRequest struct{}"
	}

	return strings.Join([]string{"UpdateAclAccountRequest", string(data)}, " ")
}
