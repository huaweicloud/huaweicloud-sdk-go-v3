package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBasicPluginResponse Response Object
type UpdateBasicPluginResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBasicPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBasicPluginResponse struct{}"
	}

	return strings.Join([]string{"UpdateBasicPluginResponse", string(data)}, " ")
}
