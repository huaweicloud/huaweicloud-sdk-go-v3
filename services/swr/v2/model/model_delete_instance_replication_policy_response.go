package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceReplicationPolicyResponse Response Object
type DeleteInstanceReplicationPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstanceReplicationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceReplicationPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceReplicationPolicyResponse", string(data)}, " ")
}
