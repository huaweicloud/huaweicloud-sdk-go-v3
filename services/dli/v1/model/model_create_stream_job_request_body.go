package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStreamJobRequestBody 提交流作业的请求参数。
type CreateStreamJobRequestBody struct {

	// 流作业事务 ID，防止重复提交。长度限制：0-64个字符。
	TransactionId string `json:"transaction_id"`

	// 作业类型：flink_sql_job或flink_jar_job
	JobType *string `json:"job_type,omitempty"`

	// 流作业描述。长度限制：0-512个字符。
	Description *string `json:"description,omitempty"`

	// 标签。
	Tags *[]Tag `json:"tags,omitempty"`

	EnvironmentConfig *StreamEnvironmentConfig `json:"environment_config,omitempty"`

	RuntimeConfig *StreamRuntimeConfig `json:"runtime_config,omitempty"`

	FlinkSqlParameter *FlinkSqlParameter `json:"flink_sql_parameter,omitempty"`

	FlinkJarParameter *FlinkJarParameter `json:"flink_jar_parameter,omitempty"`
}

func (o CreateStreamJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateStreamJobRequestBody", string(data)}, " ")
}
