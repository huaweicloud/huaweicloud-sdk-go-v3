package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWebHookConfigStatusResponse Response Object
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
