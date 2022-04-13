package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCommandRequestBody struct {
	// 服务命令名称，支持大小写字母，数字，中划线及下划线，长度2-50

	Name string `json:"name"`
	// 服务命令描述，长度0-200

	Description *string `json:"description,omitempty"`
}

func (o CreateCommandRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommandRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCommandRequestBody", string(data)}, " ")
}
