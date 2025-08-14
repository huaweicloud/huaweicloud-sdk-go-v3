package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePreBootPolicyRequest Request Object
type UpdatePreBootPolicyRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`

	Body *UpdatePreBootPolicyReq `json:"body,omitempty"`
}

func (o UpdatePreBootPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePreBootPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePreBootPolicyRequest", string(data)}, " ")
}
