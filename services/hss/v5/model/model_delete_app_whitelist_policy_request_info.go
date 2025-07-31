package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteAppWhitelistPolicyRequestInfo struct {

	// 策略ID
	PolicyId string `json:"policy_id"`
}

func (o DeleteAppWhitelistPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppWhitelistPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"DeleteAppWhitelistPolicyRequestInfo", string(data)}, " ")
}
