package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InviteShareResponse Response Object
type InviteShareResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InviteShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteShareResponse struct{}"
	}

	return strings.Join([]string{"InviteShareResponse", string(data)}, " ")
}
