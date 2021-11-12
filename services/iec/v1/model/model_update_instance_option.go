package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新边缘实例的参数
type UpdateInstanceOption struct {
	// 修改后的边缘实例名称， 只能由中文字符、英文字母、数字及“_”、“-”、“.”组成。

	Name *string `json:"name,omitempty"`
	// 描述， 不能包含“<”，“>”。

	Description *string `json:"description,omitempty"`
}

func (o UpdateInstanceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceOption struct{}"
	}

	return strings.Join([]string{"UpdateInstanceOption", string(data)}, " ")
}
