package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePluginResponse Response Object
type DeletePluginResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePluginResponse struct{}"
	}

	return strings.Join([]string{"DeletePluginResponse", string(data)}, " ")
}
