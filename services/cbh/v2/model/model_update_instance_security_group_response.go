package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceSecurityGroupResponse Response Object
type UpdateInstanceSecurityGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateInstanceSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceSecurityGroupResponse", string(data)}, " ")
}
