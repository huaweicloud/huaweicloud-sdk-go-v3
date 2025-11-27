package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateHookVersionResponse Response Object
type CreatePrivateHookVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePrivateHookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateHookVersionResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateHookVersionResponse", string(data)}, " ")
}
