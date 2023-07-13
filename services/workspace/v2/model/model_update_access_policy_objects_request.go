package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessPolicyObjectsRequest Request Object
type UpdateAccessPolicyObjectsRequest struct {

	// 接入策略id。
	AccessPolicyId string `json:"access_policy_id"`

	Body *UpdateAccessPolicyObjectsReq `json:"body,omitempty"`
}

func (o UpdateAccessPolicyObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessPolicyObjectsRequest struct{}"
	}

	return strings.Join([]string{"UpdateAccessPolicyObjectsRequest", string(data)}, " ")
}
