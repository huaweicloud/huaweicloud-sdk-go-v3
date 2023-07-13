package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetWebHookConfigResponse Response Object
type SetWebHookConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetWebHookConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetWebHookConfigResponse struct{}"
	}

	return strings.Join([]string{"SetWebHookConfigResponse", string(data)}, " ")
}
