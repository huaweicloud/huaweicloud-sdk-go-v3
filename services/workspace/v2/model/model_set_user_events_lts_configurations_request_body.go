package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetUserEventsLtsConfigurationsRequestBody 配置用户事件LTS请求体。
type SetUserEventsLtsConfigurationsRequestBody struct {

	// 是否启用日志流LTS。
	Enable bool `json:"enable"`

	// 日志组id。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流id。
	LogStreamId *string `json:"log_stream_id,omitempty"`
}

func (o SetUserEventsLtsConfigurationsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetUserEventsLtsConfigurationsRequestBody struct{}"
	}

	return strings.Join([]string{"SetUserEventsLtsConfigurationsRequestBody", string(data)}, " ")
}
