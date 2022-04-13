package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
