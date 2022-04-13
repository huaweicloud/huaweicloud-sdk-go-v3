package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求头
type ListEventItemsHeaders struct {
	// 请求长度

	ContentLength *string `json:"content-length,omitempty"`
	// 域名

	Host *string `json:"host,omitempty"`
	// 内容类型

	ContentType *string `json:"content-type,omitempty"`
	// 代理

	UserAgent *string `json:"user-agent,omitempty"`
	// 接收内容类型

	Accept *string `json:"accept,omitempty"`
}

func (o ListEventItemsHeaders) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventItemsHeaders struct{}"
	}

	return strings.Join([]string{"ListEventItemsHeaders", string(data)}, " ")
}
