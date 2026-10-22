package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBotMTrafficDetectionConditionRequest Request Object
type UpdateBotMTrafficDetectionConditionRequest struct {

	// policyId
	PolicyId string `json:"policy_id"`

	// 流量检测条件Id
	ConditionId string `json:"condition_id"`

	Body *SaveTrafficDetectionConditionRequestBody `json:"body,omitempty"`
}

func (o UpdateBotMTrafficDetectionConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMTrafficDetectionConditionRequest struct{}"
	}

	return strings.Join([]string{"UpdateBotMTrafficDetectionConditionRequest", string(data)}, " ")
}
