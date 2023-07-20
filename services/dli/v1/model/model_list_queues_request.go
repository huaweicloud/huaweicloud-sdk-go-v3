package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuesRequest Request Object
type ListQueuesRequest struct {

	// 队列的类型,。有如下三种类型： sql general all 如果不指定，默认为sql。
	QueueType *string `json:"queue_type,omitempty"`

	// 查询根据标签进行过滤
	Tags *string `json:"tags,omitempty"`

	// 是否返回收费信息
	WithChargeInfo *bool `json:"with-charge-info,omitempty"`

	// 是否返回权限信息。
	WithPriv *bool `json:"with-priv,omitempty"`
}

func (o ListQueuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuesRequest struct{}"
	}

	return strings.Join([]string{"ListQueuesRequest", string(data)}, " ")
}
