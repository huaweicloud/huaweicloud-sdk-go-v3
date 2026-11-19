package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRemindersConfigResponse Response Object
type SetRemindersConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetRemindersConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRemindersConfigResponse struct{}"
	}

	return strings.Join([]string{"SetRemindersConfigResponse", string(data)}, " ")
}
