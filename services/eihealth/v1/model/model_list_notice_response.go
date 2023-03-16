package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNoticeResponse struct {

	// 通知消息总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 通知消息未读总数
	UnreadCount *int32 `json:"unread_count,omitempty"`

	// 通知消息列表
	Notices        *[]NoticeRsp `json:"notices,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListNoticeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNoticeResponse struct{}"
	}

	return strings.Join([]string{"ListNoticeResponse", string(data)}, " ")
}
