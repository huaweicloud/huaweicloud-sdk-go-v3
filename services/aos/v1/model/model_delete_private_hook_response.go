package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateHookResponse Response Object
type DeletePrivateHookResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateHookResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateHookResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateHookResponse", string(data)}, " ")
}
