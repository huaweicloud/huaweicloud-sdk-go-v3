package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgeCloudOption 边缘业务对象
type EdgeCloudOption struct {

	// 边缘业务名称。 取值范围：只能由中文字符、大小写英文字母、数字及中划线、下划线组成，且长度为[1-32]个字符。
	Name *string `json:"name,omitempty"`

	// 已有边缘业务ID，该参数用于扩容边缘业务场景。 >-  id与name不可同时为空，同时有值时部署计划无效； - 通过id扩容场景要求区域分布层级与原边缘业务一致； - 区域分布层级为站点级的边缘业务不支持扩容。
	Id *string `json:"id,omitempty"`

	// 描述，缺省值为空字符串。
	Description *string `json:"description,omitempty"`

	Coverage *Coverage `json:"coverage"`

	Stack *Stack `json:"stack"`
}

func (o EdgeCloudOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeCloudOption struct{}"
	}

	return strings.Join([]string{"EdgeCloudOption", string(data)}, " ")
}
