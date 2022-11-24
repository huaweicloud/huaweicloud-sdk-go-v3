package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppUpdate struct {

	// 应用模板别名，中文、英文字母、数字、中划线、下划线，最大64字符
	Alias *string `json:"alias,omitempty"`

	// 应用模板描述，最大长度255，不允许^ ~ # $ % & * < > ( ) [ ] { } ' \" \\
	Description *string `json:"description,omitempty"`
}

func (o AppUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppUpdate struct{}"
	}

	return strings.Join([]string{"AppUpdate", string(data)}, " ")
}
