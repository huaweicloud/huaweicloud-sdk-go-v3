package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBotMTrafficDetectionConditionRequest Request Object
type CreateBotMTrafficDetectionConditionRequest struct {

	// policyid
	PolicyId string `json:"policy_id"`

	Body *SaveTrafficDetectionConditionRequestBody `json:"body,omitempty"`
}

func (o CreateBotMTrafficDetectionConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBotMTrafficDetectionConditionRequest struct{}"
	}

	return strings.Join([]string{"CreateBotMTrafficDetectionConditionRequest", string(data)}, " ")
}
