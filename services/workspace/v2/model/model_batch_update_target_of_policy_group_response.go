package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateTargetOfPolicyGroupResponse Response Object
type BatchUpdateTargetOfPolicyGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateTargetOfPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateTargetOfPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateTargetOfPolicyGroupResponse", string(data)}, " ")
}
