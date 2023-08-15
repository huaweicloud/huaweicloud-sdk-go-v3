package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionExternalInfo struct {

	// 插件id
	ExtensionId *string `json:"extension_id,omitempty"`

	// 源码仓地址
	RepoUrl *string `json:"repo_url,omitempty"`

	// 帮助页面
	HelpPage *string `json:"help_page,omitempty"`

	// 产品首页
	Website *string `json:"website,omitempty"`

	// 问题链接
	IssueLink *string `json:"issue_link,omitempty"`

	// 是否支持预览
	ShowPreviews *bool `json:"show_previews,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ExtensionExternalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionExternalInfo struct{}"
	}

	return strings.Join([]string{"ExtensionExternalInfo", string(data)}, " ")
}
