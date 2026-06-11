package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistorySessionsRequest Request Object
type ListHistorySessionsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 进程ID
	Pid *string `json:"pid,omitempty"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 数据库名
	DatabaseName *string `json:"database_name,omitempty"`

	// 参数解释： 开始时间。 格式为UTC时间戳。 取值范围： 不涉及。 默认取值： 不涉及。
	StartTime int64 `json:"start_time"`

	// 参数解释： 结束时间。 格式为UTC时间戳。 取值范围： 不涉及。 默认取值： 不涉及。
	EndTime int64 `json:"end_time"`

	// 参数解释： 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，值为0表示从第一条数据开始查询）。 约束限制： 必须为数字，不能为负数。 取值范围： 大于等于0的整数。 默认取值： 0
	Offset *int32 `json:"offset,omitempty"`

	// 参数解释： 查询记录数。 约束限制： 不涉及。 取值范围： [1, 1000] 默认取值： 100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHistorySessionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistorySessionsRequest struct{}"
	}

	return strings.Join([]string{"ListHistorySessionsRequest", string(data)}, " ")
}
