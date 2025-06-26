package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteInstanceReplicationPolicyResponse Response Object
type ExecuteInstanceReplicationPolicyResponse struct {

	// 执行ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExecuteInstanceReplicationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteInstanceReplicationPolicyResponse struct{}"
	}

	return strings.Join([]string{"ExecuteInstanceReplicationPolicyResponse", string(data)}, " ")
}
