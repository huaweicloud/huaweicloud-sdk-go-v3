package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserEventsLtsConfigurationsResponse Response Object
type ListUserEventsLtsConfigurationsResponse struct {

	// 是否启用日志流LTS。
	Enable *bool `json:"enable,omitempty"`

	// 日志组id。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志流id。
	LogStreamId    *string `json:"log_stream_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListUserEventsLtsConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserEventsLtsConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListUserEventsLtsConfigurationsResponse", string(data)}, " ")
}
