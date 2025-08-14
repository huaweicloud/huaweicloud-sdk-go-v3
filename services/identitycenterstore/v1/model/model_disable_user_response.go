package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableUserResponse Response Object
type DisableUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableUserResponse struct{}"
	}

	return strings.Join([]string{"DisableUserResponse", string(data)}, " ")
}
