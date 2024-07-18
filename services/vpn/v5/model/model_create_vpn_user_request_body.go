package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnUserRequestBody struct {
	User *CreateVpnUserRequestBodyContent `json:"user"`
}

func (o CreateVpnUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVpnUserRequestBody", string(data)}, " ")
}
