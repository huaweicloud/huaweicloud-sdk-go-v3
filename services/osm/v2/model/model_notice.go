package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Notice struct {

	// 公告Id
	Id string `json:"id"`

	// 公告内容
	Content string `json:"content"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o Notice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Notice struct{}"
	}

	return strings.Join([]string{"Notice", string(data)}, " ")
}
