package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceSecurityGroupResponse Response Object
type UpdateInstanceSecurityGroupResponse struct {

	// 安全组ID
	SecurityGroupId *string `json:"security_group_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o UpdateInstanceSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceSecurityGroupResponse", string(data)}, " ")
}
