package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyGroupResponse Response Object
type DeletePolicyGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyGroupResponse", string(data)}, " ")
}
