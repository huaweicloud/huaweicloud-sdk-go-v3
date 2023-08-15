package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogGroupResponse Response Object
type UpdateLogGroupResponse struct {

	// 创建该日志组的时间， 毫秒级。
	CreationTime *int64 `json:"creation_time,omitempty"`

	// 日志组的名称。
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 日志组ID。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 日志存储时间（天）。
	TtlInDays      *int32 `json:"ttl_in_days,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateLogGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogGroupResponse", string(data)}, " ")
}
