package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePluginVersionResponse Response Object
type CreatePluginVersionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePluginVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginVersionResponse struct{}"
	}

	return strings.Join([]string{"CreatePluginVersionResponse", string(data)}, " ")
}
