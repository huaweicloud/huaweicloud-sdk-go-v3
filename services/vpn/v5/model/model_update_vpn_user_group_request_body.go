package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnUserGroupRequestBody struct {
	UserGroup *UpdateVpnUserGroupRequestBodyContent `json:"user_group,omitempty"`
}

func (o UpdateVpnUserGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserGroupRequestBody", string(data)}, " ")
}
