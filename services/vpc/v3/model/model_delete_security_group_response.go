package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityGroupResponse Response Object
type DeleteSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupResponse", string(data)}, " ")
}
