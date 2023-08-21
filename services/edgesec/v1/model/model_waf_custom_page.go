package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WafCustomPage 自定义拦截页面
type WafCustomPage struct {

	// 返回状态码
	StatusCode string `json:"status_code"`

	// 页面内容类型
	ContentType string `json:"content_type"`

	// 页面内容
	Content string `json:"content"`
}

func (o WafCustomPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafCustomPage struct{}"
	}

	return strings.Join([]string{"WafCustomPage", string(data)}, " ")
}
