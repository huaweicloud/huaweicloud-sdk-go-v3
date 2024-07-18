package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnUserGroupResponse Response Object
type CreateVpnUserGroupResponse struct {
	UserGroup *CreateVpnUserGroupResponseBodyUserGroup `json:"user_group,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateVpnUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateVpnUserGroupResponse", string(data)}, " ")
}
