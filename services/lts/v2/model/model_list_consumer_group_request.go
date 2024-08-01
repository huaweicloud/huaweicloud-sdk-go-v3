package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConsumerGroupRequest Request Object
type ListConsumerGroupRequest struct {

	// 日志组ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID。 缺省值：None 最小长度：36 最大长度：36
	GroupId string `json:"group_id"`

	// 日志流ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID 缺省值：None 最小长度：36 最大长度：36
	StreamId string `json:"stream_id"`
}

func (o ListConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"ListConsumerGroupRequest", string(data)}, " ")
}
