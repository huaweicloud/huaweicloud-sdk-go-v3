package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceReplicationPolicyResponse Response Object
type UpdateInstanceReplicationPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceReplicationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceReplicationPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceReplicationPolicyResponse", string(data)}, " ")
}
