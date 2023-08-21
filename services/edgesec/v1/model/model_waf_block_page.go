package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WafBlockPage 拦截页面配置
type WafBlockPage struct {

	// 拦截模板名称
	Template string `json:"template"`

	CustomPage *WafCustomPage `json:"custom_page,omitempty"`

	// 重定向URL
	RedirectUrl *string `json:"redirect_url,omitempty"`
}

func (o WafBlockPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafBlockPage struct{}"
	}

	return strings.Join([]string{"WafBlockPage", string(data)}, " ")
}
