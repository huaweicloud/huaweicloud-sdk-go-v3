package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateLogGroupResponse struct {

	// 创建该日志组的时间， 毫秒级。
	CreationTime *int64 `json:"creation_time,omitempty" xml:"creation_time"`

	// 日志组的名称。
	LogGroupName *string `json:"log_group_name,omitempty" xml:"log_group_name"`

	// 日志组ID。
	LogGroupId *string `json:"log_group_id,omitempty" xml:"log_group_id"`

	// 日志存储时间（天）。
	TtlInDays      *int32 `json:"ttl_in_days,omitempty" xml:"ttl_in_days"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateLogGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogGroupResponse", string(data)}, " ")
}
