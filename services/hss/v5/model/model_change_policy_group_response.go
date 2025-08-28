package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePolicyGroupResponse Response Object
type ChangePolicyGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangePolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ChangePolicyGroupResponse", string(data)}, " ")
}
