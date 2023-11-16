package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginOutputsResponse Response Object
type ShowPluginOutputsResponse struct {
	Body           *[]PluginPartQueryVoListAgentPluginOutputVo `json:"body,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o ShowPluginOutputsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginOutputsResponse struct{}"
	}

	return strings.Join([]string{"ShowPluginOutputsResponse", string(data)}, " ")
}
