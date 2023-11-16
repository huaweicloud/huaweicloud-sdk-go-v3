package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBasicPluginResponse Response Object
type CreateBasicPluginResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBasicPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBasicPluginResponse struct{}"
	}

	return strings.Join([]string{"CreateBasicPluginResponse", string(data)}, " ")
}
