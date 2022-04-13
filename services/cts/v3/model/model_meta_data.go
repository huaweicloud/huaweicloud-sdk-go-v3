package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标识为扩展字段，包括count（当前响应中事件记录的个数）和marker（本次查询返回事件列表最后一个事件ID）。
type MetaData struct {
	// 标识本次查询事件列表返回的事件记录的总条数。

	Count *int32 `json:"count,omitempty"`
	// 标识本次查询事件列表返回的最后一个事件ID。可以使用这个参数返回值作为分页请求参数next的值，如果marker返回为null，则表示当前请求条件下查询事件列表已经全部返回没有下一页。

	Marker *string `json:"marker,omitempty"`
}

func (o MetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaData struct{}"
	}

	return strings.Join([]string{"MetaData", string(data)}, " ")
}
