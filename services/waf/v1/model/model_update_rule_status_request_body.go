package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleStatusRequestBody struct {
	// 状态（开启：1，关闭：0）

	Status *int32 `json:"status,omitempty"`
}

func (o UpdateRuleStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRuleStatusRequestBody", string(data)}, " ")
}
