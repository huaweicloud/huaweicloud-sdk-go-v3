package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteAccessPoliciesReq struct {

	// 策略ID列表。
	PolicyIdList *[]string `json:"policy_id_list,omitempty"`
}

func (o BatchDeleteAccessPoliciesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAccessPoliciesReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteAccessPoliciesReq", string(data)}, " ")
}
