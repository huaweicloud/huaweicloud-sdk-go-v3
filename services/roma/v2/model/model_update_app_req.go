package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAppReq struct {
	// 应用名称 - 字符集：支持中文、英文字母、数字、中划线、下划线、点、空格和中英文圆括号 - 约束：实例下唯一

	Name *string `json:"name,omitempty"`
	// 应用描述

	Remark *string `json:"remark,omitempty"`
	// 是否收藏应用，收藏的应用会在列表里优先显示

	Favorite *bool `json:"favorite,omitempty"`
}

func (o UpdateAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppReq struct{}"
	}

	return strings.Join([]string{"UpdateAppReq", string(data)}, " ")
}
