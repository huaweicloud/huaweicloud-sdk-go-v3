package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchRaspResponse Response Object
type SwitchRaspResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchRaspResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchRaspResponse struct{}"
	}

	return strings.Join([]string{"SwitchRaspResponse", string(data)}, " ")
}
