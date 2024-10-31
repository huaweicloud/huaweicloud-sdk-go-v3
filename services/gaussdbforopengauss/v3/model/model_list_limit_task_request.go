package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLimitTaskRequest Request Object
type ListLimitTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 限流任务范围，目前支持SQL,SESSION。
	TaskScope *string `json:"task_scope,omitempty"`

	// 限流类型，支持SQL_ID、SQL_TYPE、SESSION_ACTIVE_MAX_COUNT类型。
	LimitType *string `json:"limit_type,omitempty"`

	// 限流类型值，支持模糊匹配。
	LimitTypeValue *string `json:"limit_type_value,omitempty"`

	// 限流任务名，支持模糊匹配。
	TaskName *string `json:"task_name,omitempty"`

	// sql模板，支持模糊匹配。
	SqlModel *string `json:"sql_model,omitempty"`

	// 规则名。
	RuleName *string `json:"rule_name,omitempty"`

	// 限流任务开始时间，格式为yyyy-mm-ddThh:mm:ssZ,当前时间指UTC时间。
	StartTime *string `json:"start_time,omitempty"`

	// 限流任务结束时间，格式为yyyy-mm-ddThh:mm:ssZ,当前时间指UTC时间。
	EndTime *string `json:"end_time,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。  取值范围：0 - 10000
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"ListLimitTaskRequest", string(data)}, " ")
}
