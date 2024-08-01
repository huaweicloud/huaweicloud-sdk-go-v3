package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCursorByTimeRequest Request Object
type ShowCursorByTimeRequest struct {

	// 日志组ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID。 缺省值：None 最小长度：36 最大长度：36
	GroupId string `json:"group_id"`

	// 日志流ID，获取方式请参见：获取项目ID，获取账号ID，日志组ID、日志流ID 缺省值：None 最小长度：36 最大长度：36
	StreamId string `json:"stream_id"`

	// Shrad ID
	ShardId string `json:"shard_id"`

	// 起始时间戳，时间单位为纳秒
	From string `json:"from"`
}

func (o ShowCursorByTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCursorByTimeRequest struct{}"
	}

	return strings.Join([]string{"ShowCursorByTimeRequest", string(data)}, " ")
}
