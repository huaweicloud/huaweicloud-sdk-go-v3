package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateWebHookConfigStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWebHookConfigStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebHookConfigStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateWebHookConfigStatusResponse", string(data)}, " ")
}
