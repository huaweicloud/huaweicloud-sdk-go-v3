package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronDeleteSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteSecurityGroupResponse", string(data)}, " ")
}
