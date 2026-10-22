package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiAppPolicyUpdateRequest struct {

	// 要与应用关联的策略ID
	PolicyId string `json:"policy_id"`
}

func (o AiAppPolicyUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiAppPolicyUpdateRequest struct{}"
	}

	return strings.Join([]string{"AiAppPolicyUpdateRequest", string(data)}, " ")
}
