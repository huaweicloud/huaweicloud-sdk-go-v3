package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityGroupTagResponse Response Object
type DeleteSecurityGroupTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityGroupTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupTagResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupTagResponse", string(data)}, " ")
}
