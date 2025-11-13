package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResolverQueryLogConfig struct {

	// 解析器访问日志的ID，UUID形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// 日志组ID。
	LtsGroupId *string `json:"lts_group_id,omitempty"`

	// 日志流ID。
	LtsTopicId *string `json:"lts_topic_id,omitempty"`

	// 关联的VPC ID列表。
	VpcIds *[]string `json:"vpc_ids,omitempty"`
}

func (o ResolverQueryLogConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResolverQueryLogConfig struct{}"
	}

	return strings.Join([]string{"ResolverQueryLogConfig", string(data)}, " ")
}
