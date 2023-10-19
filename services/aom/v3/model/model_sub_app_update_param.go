package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubAppUpdateParam struct {

	// 子应用名称：字符集长度2-64，仅支持字符集：英文字母、数字、下划线、中划线、点
	Name string `json:"name"`

	// 子应用节点显示名：字符集长度2-64，仅支持字符集：中文字符、英文字母、数字、下划线、中划线、点
	DisplayName *string `json:"display_name,omitempty"`

	// 描述：最大255字符
	Description *string `json:"description,omitempty"`
}

func (o SubAppUpdateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubAppUpdateParam struct{}"
	}

	return strings.Join([]string{"SubAppUpdateParam", string(data)}, " ")
}
