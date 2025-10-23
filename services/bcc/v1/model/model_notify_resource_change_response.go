package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotifyResourceChangeResponse Response Object
type NotifyResourceChangeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NotifyResourceChangeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifyResourceChangeResponse struct{}"
	}

	return strings.Join([]string{"NotifyResourceChangeResponse", string(data)}, " ")
}
