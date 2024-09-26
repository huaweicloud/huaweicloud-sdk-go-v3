package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHttpPolicyRuleStatusRequestBody struct {

	// 0-关闭 1-开启
	Status int32 `json:"status"`
}

func (o UpdateHttpPolicyRuleStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpPolicyRuleStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHttpPolicyRuleStatusRequestBody", string(data)}, " ")
}
