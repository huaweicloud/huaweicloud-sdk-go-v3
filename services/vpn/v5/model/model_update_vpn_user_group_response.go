package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnUserGroupResponse Response Object
type UpdateVpnUserGroupResponse struct {
	UserGroup *VpnUserGroup `json:"user_group,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateVpnUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserGroupResponse", string(data)}, " ")
}
