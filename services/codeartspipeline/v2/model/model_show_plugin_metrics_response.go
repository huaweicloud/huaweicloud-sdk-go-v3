package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginMetricsResponse Response Object
type ShowPluginMetricsResponse struct {
	Body           *[]PluginPartQueryVoListAgentPluginOutputVo `json:"body,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o ShowPluginMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowPluginMetricsResponse", string(data)}, " ")
}
