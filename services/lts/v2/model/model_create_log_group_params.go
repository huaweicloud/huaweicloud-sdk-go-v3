package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLogGroupParams 创建日志组参数。
type CreateLogGroupParams struct {

	// 需要创建的日志组名称。
	LogGroupName string `json:"log_group_name"`

	// 日志存储时间（天），取值范围：1-30。
	TtlInDays int32 `json:"ttl_in_days"`

	// 标签字段信息
	Tags *[]TagsBody `json:"tags,omitempty"`

	// 日志组别名
	LogGroupNameAlias *string `json:"log_group_name_alias,omitempty"`
}

func (o CreateLogGroupParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLogGroupParams struct{}"
	}

	return strings.Join([]string{"CreateLogGroupParams", string(data)}, " ")
}
