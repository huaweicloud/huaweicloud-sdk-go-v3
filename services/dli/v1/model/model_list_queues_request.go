package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuesRequest Request Object
type ListQueuesRequest struct {

	// 参数解释:  队列的类型 示例: sql 约束限制:  无 取值范围: sql, general, all 默认取值: sql
	QueueType *string `json:"queue_type,omitempty"`

	// 参数解释: 查询根据标签进行过滤 示例: taga=tagb,owner=ph 约束限制:  符合“key1=value1,key2=value2”的字符串 取值范围: 无 默认取值: 无
	Tags *string `json:"tags,omitempty"`

	// 是否返回收费信息 示例: true 约束限制:  无 取值范围: true, false 默认取值: 无
	WithChargeInfo *bool `json:"with-charge-info,omitempty"`

	// 是否返回权限信息 示例: true 约束限制:  无 取值范围: true, false 默认取值: 无
	WithPriv *bool `json:"with-priv,omitempty"`
}

func (o ListQueuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuesRequest struct{}"
	}

	return strings.Join([]string{"ListQueuesRequest", string(data)}, " ")
}
