package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateHookVersionResponse Response Object
type DeletePrivateHookVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateHookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateHookVersionResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateHookVersionResponse", string(data)}, " ")
}
