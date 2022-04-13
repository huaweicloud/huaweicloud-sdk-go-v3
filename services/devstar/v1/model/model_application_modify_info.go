package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationModifyInfo struct {
	// 应用名称

	Name *string `json:"name,omitempty"`
	// 应用描述

	Description *string `json:"description,omitempty"`
	// 应用图标(传入图片的Base64编码,大小限制15k)

	Icon *string `json:"icon,omitempty"`
}

func (o ApplicationModifyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationModifyInfo struct{}"
	}

	return strings.Join([]string{"ApplicationModifyInfo", string(data)}, " ")
}
