package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopPluginResponse Response Object
type StopPluginResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopPluginResponse struct{}"
	}

	return strings.Join([]string{"StopPluginResponse", string(data)}, " ")
}
