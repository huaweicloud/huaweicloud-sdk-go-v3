package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNoticeRequest Request Object
type ListNoticeRequest struct {

	// 消息状态是否已读，true返回已读，false返回未读，不填返回全部
	IsRead *bool `json:"is_read,omitempty"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 查询偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListNoticeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNoticeRequest struct{}"
	}

	return strings.Join([]string{"ListNoticeRequest", string(data)}, " ")
}
