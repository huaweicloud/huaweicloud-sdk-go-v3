package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAiGuardPolicyResponse Response Object
type BatchDeleteAiGuardPolicyResponse struct {

	// 已删除的策略 ID 列表
	PolicyIds      *[]string `json:"policy_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteAiGuardPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAiGuardPolicyResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAiGuardPolicyResponse", string(data)}, " ")
}
