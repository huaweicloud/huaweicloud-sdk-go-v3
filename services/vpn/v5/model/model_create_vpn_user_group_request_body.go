package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnUserGroupRequestBody struct {
	UserGroup *CreateVpnUserGroupRequestBodyContent `json:"user_group"`
}

func (o CreateVpnUserGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpnUserGroupRequestBody", string(data)}, " ")
}
