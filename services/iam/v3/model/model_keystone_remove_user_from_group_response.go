package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneRemoveUserFromGroupResponse Response Object
type KeystoneRemoveUserFromGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneRemoveUserFromGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneRemoveUserFromGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneRemoveUserFromGroupResponse", string(data)}, " ")
}
