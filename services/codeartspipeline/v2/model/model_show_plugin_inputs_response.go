package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginInputsResponse Response Object
type ShowPluginInputsResponse struct {
	Body           *[]PluginPartQueryVoListAgentPluginInputVo `json:"body,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o ShowPluginInputsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginInputsResponse struct{}"
	}

	return strings.Join([]string{"ShowPluginInputsResponse", string(data)}, " ")
}
