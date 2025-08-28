package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddPolicyGroupResponse Response Object
type AddPolicyGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"AddPolicyGroupResponse", string(data)}, " ")
}
