package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceReplicationPolicyResponse Response Object
type CreateInstanceReplicationPolicyResponse struct {

	// 命名空间ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateInstanceReplicationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceReplicationPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceReplicationPolicyResponse", string(data)}, " ")
}
