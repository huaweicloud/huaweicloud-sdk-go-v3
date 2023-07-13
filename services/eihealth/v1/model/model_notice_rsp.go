package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NoticeRsp 消息通知返回体
type NoticeRsp struct {

	// 消息id
	Id *string `json:"id,omitempty"`

	// 消息类型
	Type *string `json:"type,omitempty"`

	// 消息详情
	Detail *string `json:"detail,omitempty"`

	// 消息创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 是否已读
	IsRead *bool `json:"is_read,omitempty"`
}

func (o NoticeRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoticeRsp struct{}"
	}

	return strings.Join([]string{"NoticeRsp", string(data)}, " ")
}
