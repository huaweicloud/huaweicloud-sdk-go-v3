package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自定义告警页面
type CustomPage struct {

	// 返回状态码
	StatusCode string `json:"status_code"`

	// “自定义”告警页面内容类型，可选择text/html、text/xml和application/json三种类型
	ContentType string `json:"content_type"`

	// 根据选择的“页面类型”配置对应的页面内容，具体示例可以参考“Web应用防火墙 WAF”用户手册
	Content string `json:"content"`
}

func (o CustomPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomPage struct{}"
	}

	return strings.Join([]string{"CustomPage", string(data)}, " ")
}
