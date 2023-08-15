package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneDeleteGroupResponse Response Object
type KeystoneDeleteGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneDeleteGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneDeleteGroupResponse struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteGroupResponse", string(data)}, " ")
}
