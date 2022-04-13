package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowL7PolicyRequest struct {
	// 转发策略ID。

	L7policyId string `json:"l7policy_id"`
}

func (o ShowL7PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowL7PolicyRequest", string(data)}, " ")
}
