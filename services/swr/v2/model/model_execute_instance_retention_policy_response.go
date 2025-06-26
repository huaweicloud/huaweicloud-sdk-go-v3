package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteInstanceRetentionPolicyResponse Response Object
type ExecuteInstanceRetentionPolicyResponse struct {

	// 老化策略执行ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExecuteInstanceRetentionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteInstanceRetentionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ExecuteInstanceRetentionPolicyResponse", string(data)}, " ")
}
