package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOnlineConfAttendeeResponse Response Object
type ListOnlineConfAttendeeResponse struct {

	// 在线与会者信息列表
	Data *[]OnlineAttendeeRecordInfo `json:"data,omitempty"`

	// 记录数偏移,第几条
	Offset *int32 `json:"offset,omitempty"`

	// 每页的记录数
	Limit *int32 `json:"limit,omitempty"`

	// 总记录数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOnlineConfAttendeeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOnlineConfAttendeeResponse struct{}"
	}

	return strings.Join([]string{"ListOnlineConfAttendeeResponse", string(data)}, " ")
}
