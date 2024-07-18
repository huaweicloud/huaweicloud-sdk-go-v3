package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnUserGroupResponse Response Object
type ShowVpnUserGroupResponse struct {
	UserGroup *VpnUserGroup `json:"user_group,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowVpnUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnUserGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowVpnUserGroupResponse", string(data)}, " ")
}
