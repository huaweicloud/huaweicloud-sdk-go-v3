package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateReplicationExecutionRequestBody struct {

	// 策略ID
	PolicyId int32 `json:"policy_id"`
}

func (o CreateReplicationExecutionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReplicationExecutionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateReplicationExecutionRequestBody", string(data)}, " ")
}
