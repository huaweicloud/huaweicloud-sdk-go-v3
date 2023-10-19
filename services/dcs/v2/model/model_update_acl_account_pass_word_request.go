package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclAccountPassWordRequest Request Object
type UpdateAclAccountPassWordRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// ACL账号ID。
	AccountId string `json:"account_id"`

	Body *AclAccountModifyPasswordBody `json:"body,omitempty"`
}

func (o UpdateAclAccountPassWordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclAccountPassWordRequest struct{}"
	}

	return strings.Join([]string{"UpdateAclAccountPassWordRequest", string(data)}, " ")
}
