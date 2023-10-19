package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclAccountRemarkRequest Request Object
type UpdateAclAccountRemarkRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// ACL账号ID。
	AccountId string `json:"account_id"`

	Body *UpdateAclAccountRemarkRequestBody `json:"body,omitempty"`
}

func (o UpdateAclAccountRemarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclAccountRemarkRequest struct{}"
	}

	return strings.Join([]string{"UpdateAclAccountRemarkRequest", string(data)}, " ")
}
