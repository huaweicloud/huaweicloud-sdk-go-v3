package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePolicyGroupReq 创建策略组请求。
type CreatePolicyGroupReq struct {
	PolicyGroup *PolicyGroupForCreate `json:"policy_group"`
}

func (o CreatePolicyGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyGroupReq struct{}"
	}

	return strings.Join([]string{"CreatePolicyGroupReq", string(data)}, " ")
}
