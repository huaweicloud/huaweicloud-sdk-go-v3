package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRaspSwitchResponse Response Object
type SetRaspSwitchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetRaspSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRaspSwitchResponse struct{}"
	}

	return strings.Join([]string{"SetRaspSwitchResponse", string(data)}, " ")
}
