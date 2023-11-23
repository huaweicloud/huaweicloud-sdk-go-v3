package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppTemplateResourceDetail struct {

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 图标文件（base64编码）
	Icon *string `json:"icon,omitempty"`

	// 超链接地址
	Href *string `json:"href,omitempty"`
}

func (o AppTemplateResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppTemplateResourceDetail struct{}"
	}

	return strings.Join([]string{"AppTemplateResourceDetail", string(data)}, " ")
}
