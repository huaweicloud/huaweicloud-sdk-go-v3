package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchModeResponse Response Object
type SwitchModeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchModeResponse struct{}"
	}

	return strings.Join([]string{"SwitchModeResponse", string(data)}, " ")
}
