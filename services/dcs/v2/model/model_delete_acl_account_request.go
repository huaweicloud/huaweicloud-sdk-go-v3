package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAclAccountRequest Request Object
type DeleteAclAccountRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// ACL账号ID。
	AccountId string `json:"account_id"`
}

func (o DeleteAclAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAclAccountRequest struct{}"
	}

	return strings.Join([]string{"DeleteAclAccountRequest", string(data)}, " ")
}
