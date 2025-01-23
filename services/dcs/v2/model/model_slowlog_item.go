package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowlogItem 慢日志单个条目
type SlowlogItem struct {

	// 慢日志的唯一标识
	Id *int32 `json:"id,omitempty"`

	// 慢命令
	Command *string `json:"command,omitempty"`

	// 执行开始时间,格式为“2020-06-19T07:06:07Z”
	StartTime *string `json:"start_time,omitempty"`

	// 持续时间，单位是ms
	Duration *string `json:"duration,omitempty"`

	// 慢命令所在的分片名称，仅在实例类型为集群时支持
	ShardName *string `json:"shard_name,omitempty"`

	// 数据库id，当前只对指定客户开放
	DatabaseId *int32 `json:"database_id,omitempty"`

	// 操作慢日志的账号名称，当前只对指定客户开放
	Username *string `json:"username,omitempty"`

	// **参数解释**： 节点类型。 **取值范围**： 不涉及。
	NodeRole *string `json:"node_role,omitempty"`

	// **参数解释**： 客户端IP地址。 **取值范围**： 不涉及。
	ClientIp *string `json:"client_ip,omitempty"`
}

func (o SlowlogItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogItem struct{}"
	}

	return strings.Join([]string{"SlowlogItem", string(data)}, " ")
}
