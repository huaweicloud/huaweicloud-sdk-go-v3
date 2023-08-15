package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaDisassociateSecurityGroupResponse Response Object
type NovaDisassociateSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NovaDisassociateSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaDisassociateSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"NovaDisassociateSecurityGroupResponse", string(data)}, " ")
}
