package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAppReq struct {
	// 应用名称 - 字符集：支持中文、英文字母、数字、中划线、下划线、点、空格和中英文圆括号 - 约束：实例下唯一

	Name string `json:"name"`
	// 应用描述

	Remark *string `json:"remark,omitempty"`
	// 应用认证访问KEY,未提供时随机生成 - 字符集：支持中文、英文字母、数字、中划线、下划线、@号和点，以字母或中文或数字开头 - 约束：实例下唯一

	Key *string `json:"key,omitempty"`
	// 应用认证访问SECRET,未提供（字段不存在或值为null）时随机生成 - 字符集：英文字母、数字、！、@、#、$、%、+、=、点、中划线、斜线/ - 复杂度：满足大小写字母、数字、特殊字符的复杂度组合，考虑兼容性暂时可不做

	Secret *string `json:"secret,omitempty"`
	// 是否收藏应用，收藏的应用会在列表里优先显示

	Favorite *bool `json:"favorite,omitempty"`
}

func (o CreateAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppReq struct{}"
	}

	return strings.Join([]string{"CreateAppReq", string(data)}, " ")
}
