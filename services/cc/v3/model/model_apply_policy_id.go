package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyPolicyId struct {

	// 应用策略ID。
	ApplyPolicyId string `json:"apply_policy_id"`
}

func (o ApplyPolicyId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyPolicyId struct{}"
	}

	return strings.Join([]string{"ApplyPolicyId", string(data)}, " ")
}
