package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ValidateRomaAppResponse struct {
	// 应用ID

	Id *string `json:"id,omitempty"`
	// 应用名称 - 字符集：支持中文、英文字母、数字、中划线、下划线、点、空格和中英文圆括号 - 约束：实例下唯一

	Name *string `json:"name,omitempty"`
	// 应用描述

	Remark         *string `json:"remark,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateRomaAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateRomaAppResponse struct{}"
	}

	return strings.Join([]string{"ValidateRomaAppResponse", string(data)}, " ")
}
