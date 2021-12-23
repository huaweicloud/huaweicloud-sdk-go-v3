package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自定义告警页面
type CustomPage struct {
	// 返回状态码

	StatusCode string `json:"status_code"`
	// 页面内容类型

	ContentType string `json:"content_type"`
	// 页面内容

	Content string `json:"content"`
}

func (o CustomPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomPage struct{}"
	}

	return strings.Join([]string{"CustomPage", string(data)}, " ")
}
