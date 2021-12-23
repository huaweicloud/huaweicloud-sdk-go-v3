package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 告警页面
type BlockPage struct {
	// 模板名称

	Template string `json:"template"`

	CustomPage *CustomPage `json:"custom_page,omitempty"`
	// 重定向URL

	RedirectUrl *string `json:"redirect_url,omitempty"`
}

func (o BlockPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlockPage struct{}"
	}

	return strings.Join([]string{"BlockPage", string(data)}, " ")
}
