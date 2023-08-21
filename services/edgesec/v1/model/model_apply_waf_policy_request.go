package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyWafPolicyRequest Request Object
type ApplyWafPolicyRequest struct {

	// 防护策略id，t通过查询策略列表（ListPolicy）获取
	PolicyId string `json:"policy_id"`

	Body *ApplyWafPolicyRequestBody `json:"body,omitempty"`
}

func (o ApplyWafPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyWafPolicyRequest struct{}"
	}

	return strings.Join([]string{"ApplyWafPolicyRequest", string(data)}, " ")
}
