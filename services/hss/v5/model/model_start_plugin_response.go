package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartPluginResponse Response Object
type StartPluginResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPluginResponse struct{}"
	}

	return strings.Join([]string{"StartPluginResponse", string(data)}, " ")
}
