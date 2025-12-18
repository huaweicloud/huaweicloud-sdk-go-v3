package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotClusterConfigResponse Response Object
type ShowAutopilotClusterConfigResponse struct {

	// 存储时长
	TtlInDays *int32 `json:"ttl_in_days,omitempty"`

	// 日志配置项
	LogConfigs     *[]AutopilotClusterLogConfigLogConfigs `json:"log_configs,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ShowAutopilotClusterConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotClusterConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAutopilotClusterConfigResponse", string(data)}, " ")
}
