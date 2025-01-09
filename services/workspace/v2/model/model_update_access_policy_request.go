package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessPolicyRequest Request Object
type UpdateAccessPolicyRequest struct {

	// 接入策略id。
	AccessPolicyId string `json:"access_policy_id"`

	Body *UpdateAccessPolicyReq `json:"body,omitempty"`
}

func (o UpdateAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAccessPolicyRequest", string(data)}, " ")
}
