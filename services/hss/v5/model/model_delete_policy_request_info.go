package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyRequestInfo 删除策略请求
type DeletePolicyRequestInfo struct {

	// 策略名称
	PolicyName string `json:"policy_name"`

	// 策略ID
	PolicyId string `json:"policy_id"`
}

func (o DeletePolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"DeletePolicyRequestInfo", string(data)}, " ")
}
