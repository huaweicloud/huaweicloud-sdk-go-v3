package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetAclAccountPassWordRequest Request Object
type ResetAclAccountPassWordRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// ACL账号ID。
	AccountId string `json:"account_id"`

	Body *AclAccountResetPasswordBody `json:"body,omitempty"`
}

func (o ResetAclAccountPassWordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetAclAccountPassWordRequest struct{}"
	}

	return strings.Join([]string{"ResetAclAccountPassWordRequest", string(data)}, " ")
}
