package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyHttpPolicyRequest Request Object
type ApplyHttpPolicyRequest struct {

	// 防护策略id
	PolicyId string `json:"policy_id"`

	Body *ApplyHttpPolicyRequestBody `json:"body,omitempty"`
}

func (o ApplyHttpPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyHttpPolicyRequest struct{}"
	}

	return strings.Join([]string{"ApplyHttpPolicyRequest", string(data)}, " ")
}
