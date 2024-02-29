package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogStreamParams 创建日志流参数。
type CreateLogStreamParams struct {

	// 需要创建的日志流名称。
	LogStreamName string `json:"log_stream_name"`

	// 日志存储时间 说明： 该参数仅对华东-上海一、华北-北京四、华南-广州用户开放。
	TtlInDays *int32 `json:"ttl_in_days,omitempty"`

	// 标签字段信息
	Tags *[]TagsBody `json:"tags,omitempty"`

	// 日志流名称别名
	LogStreamNameAlias *string `json:"log_stream_name_alias,omitempty"`

	// 企业项目名称 >只能由中文、英文字母、数字、下划线、中划线组成，且不能使用任何大小写形式的“default”； 描述不超过512个字符。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
}

func (o CreateLogStreamParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogStreamParams struct{}"
	}

	return strings.Join([]string{"CreateLogStreamParams", string(data)}, " ")
}
