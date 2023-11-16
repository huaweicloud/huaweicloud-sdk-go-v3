package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishPluginBindResponse Response Object
type PublishPluginBindResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishPluginBindResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishPluginBindResponse struct{}"
	}

	return strings.Join([]string{"PublishPluginBindResponse", string(data)}, " ")
}
