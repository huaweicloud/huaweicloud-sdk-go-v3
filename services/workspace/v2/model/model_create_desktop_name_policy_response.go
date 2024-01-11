package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopNamePolicyResponse Response Object
type CreateDesktopNamePolicyResponse struct {

	// 策略id。
	PolicyId       *string `json:"policy_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDesktopNamePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopNamePolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateDesktopNamePolicyResponse", string(data)}, " ")
}
