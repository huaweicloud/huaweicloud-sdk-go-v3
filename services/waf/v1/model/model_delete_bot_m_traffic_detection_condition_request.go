package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBotMTrafficDetectionConditionRequest Request Object
type DeleteBotMTrafficDetectionConditionRequest struct {

	// policyId
	PolicyId string `json:"policy_id"`

	// 流量检测条件Id
	ConditionId string `json:"condition_id"`
}

func (o DeleteBotMTrafficDetectionConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBotMTrafficDetectionConditionRequest struct{}"
	}

	return strings.Join([]string{"DeleteBotMTrafficDetectionConditionRequest", string(data)}, " ")
}
