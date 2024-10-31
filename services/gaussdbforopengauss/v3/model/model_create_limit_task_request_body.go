package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLimitTaskRequestBody struct {

	// 限流任务范围，目前支持SQL,SESSION级别范围。
	TaskScope string `json:"task_scope"`

	// 任务开始时间,取值范围：非空且大于等于当前时间的前俩分钟，格式必须为yyyy-mm-ddThh:mm:ssZ,当前时间指UTC时间，SQL范围必传。
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间,取值范围：非空且大于任务开始时间，格式必须为yyyy-mm-ddThh:mm:ssZ,当前时间指UTC时间，SQL范围必传。
	EndTime *string `json:"end_time,omitempty"`

	// 限流类型：当限流范围为SQL级别时，可选SQL_ID、SQL_TYPE类型，当限流范围为SESSION级别时，可选SESSION_ACTIVE_MAX_COUNT类型。
	LimitType string `json:"limit_type"`

	// 限流类型，当限流类型为SQL_ID类型时，该值为选中模板的sqlId，当限流类型为SQL_TYPE类型时，值为SQL类型，目前支持select，update，insert，delete，meger五种值，当限流类型为SESSION_ACTIVE_MAX_COUNT类型时，只支持CPU_OR_MEMORY一种值。
	LimitTypeValue string `json:"limit_type_value"`

	// 关键词，当且仅当类型为SQL_TYPE，必传，多个关键词以逗号隔开，数量范围为[2，100]个，每个关键词长度范围[2，64]位，关键词不允许包含 \" 或 \\ 或 {} 或 null值 以及非首尾的空格符。
	KeyWords *string `json:"key_words,omitempty"`

	// 限流任务名，必传，只能为英文字母大小写，下划线，数字和$符，最大长度为100个字符。
	TaskName string `json:"task_name"`

	// CN节点执行过的SQL模板,如果类型为SQLID，则为必传。
	SqlModel *string `json:"sql_model,omitempty"`

	// 并发数，大于等于零的正整数，取值范围[0, 2147483647]。
	ParallelSize int32 `json:"parallel_size"`

	// cpu利用率阈值，正整数，取值范围[0,100）,如果类型为SESSION_ACTIVE_MAX_COUNT，必传，不支持和内存利用率阈值同时为0，如果选择只限制CPU、内存中的其中一个，则另一个必须传值0。
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// 内存利用率阈值，正整数，取值范围[0,100）,如果类型为SESSION_ACTIVE_MAX_COUNT，必传，不支持和cpu利用率阈值同时为0，如果选择只限制CPU、内存中的其中一个，则另一个必须传值0。
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`

	// CN节点数据库组,每个数据库字符串以逗号形式隔,如果类型为SQL_TYPE，则为必传。
	Databases *string `json:"databases,omitempty"`

	// CN节点信息列表，如果类型为SQL_ID，则为必传
	NodeInfos *[]CreateLimitTaskNodeOption `json:"node_infos,omitempty"`
}

func (o CreateLimitTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLimitTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateLimitTaskRequestBody", string(data)}, " ")
}
