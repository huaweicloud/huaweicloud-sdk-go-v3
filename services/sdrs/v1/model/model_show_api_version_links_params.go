package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// API的URL地址。
type ShowApiVersionLinksParams struct {
	// 链接的描述

	Rel string `json:"rel"`
	// 版本号查询链接

	Href string `json:"href"`
}

func (o ShowApiVersionLinksParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionLinksParams struct{}"
	}

	return strings.Join([]string{"ShowApiVersionLinksParams", string(data)}, " ")
}
