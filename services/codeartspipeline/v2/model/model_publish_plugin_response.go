package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishPluginResponse Response Object
type PublishPluginResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishPluginResponse struct{}"
	}

	return strings.Join([]string{"PublishPluginResponse", string(data)}, " ")
}
