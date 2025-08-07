package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnUserGroupResponseBodyUserGroup VPN用户组
type CreateVpnUserGroupResponseBodyUserGroup struct {

	// VPN用户组ID
	Id *string `json:"id,omitempty"`
}

func (o CreateVpnUserGroupResponseBodyUserGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserGroupResponseBodyUserGroup struct{}"
	}

	return strings.Join([]string{"CreateVpnUserGroupResponseBodyUserGroup", string(data)}, " ")
}
