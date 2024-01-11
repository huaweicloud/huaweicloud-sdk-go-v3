package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopNamePolicyRequest Request Object
type UpdateDesktopNamePolicyRequest struct {

	// 策略ID。
	PolicyId string `json:"policy_id"`

	Body *UpdateDesktopNamePolicyReq `json:"body,omitempty"`
}

func (o UpdateDesktopNamePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopNamePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDesktopNamePolicyRequest", string(data)}, " ")
}
