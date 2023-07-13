package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BlockPage 告警页面配置参数，用于配置“自定义”或者“重定向”页面
type BlockPage struct {

	// 模板名称
	Template string `json:"template"`

	CustomPage *CustomPage `json:"custom_page,omitempty"`

	// “重定向”页面URL
	RedirectUrl *string `json:"redirect_url,omitempty"`
}

func (o BlockPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlockPage struct{}"
	}

	return strings.Join([]string{"BlockPage", string(data)}, " ")
}
