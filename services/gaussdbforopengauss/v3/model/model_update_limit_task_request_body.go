package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLimitTaskRequestBody struct {

	// 任务开始时间，格式为yyyy-mm-ddThh:mm:ssZ，如果存在更改，需大于当前时间前俩分钟，当前时间指UTC时，SQL范围必传。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间,格式为yyyy-mm-ddThh:mm:ssZ，大于任务开始时间，当前时间指UTC时间，SQL范围必传。
	EndTime *string `json:"end_time,omitempty"`

	// 关键词，当且仅当类型为SQL_TYPE，必传，多个关键词以逗号隔开，数量范围为[2，100]个，每个关键词长度范围[2，64]位，关键词不允许包含 \" 或 \\ 或 {} 或 null值 以及非首尾的空格符。
	KeyWords *string `json:"key_words,omitempty"`

	// 并发数，大于等于零的正整数，取值范围[0, 2147483647]。
	ParallelSize *int32 `json:"parallel_size,omitempty"`

	// 限流任务名，只能为英文字母大小写，下划线，数字和$符，最大长度为100个字符。
	TaskName *string `json:"task_name,omitempty"`

	// cpu利用率阈值，正整数，取值范围[0,100）,如果类型为SESSION_ACTIVE_MAX_COUNT，必传，不支持和内存利用率阈值同时为0，如果选择只限制CPU、内存中的其中一个，则另一个必须传值0。
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// 内存利用率阈值，正整数，取值范围[0,100）,如果类型为SESSION_ACTIVE_MAX_COUNT，必传，不支持和cpu利用率阈值同时为0，如果选择只限制CPU、内存中的其中一个，则另一个必须传值0。
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`

	// CN节点数据库组，每个数据库字符串以逗号形式隔开。
	Databases *string `json:"databases,omitempty"`
}

func (o UpdateLimitTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLimitTaskRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLimitTaskRequestBody", string(data)}, " ")
}
