package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type InviteUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InviteUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteUserResponse struct{}"
	}

	return strings.Join([]string{"InviteUserResponse", string(data)}, " ")
}
