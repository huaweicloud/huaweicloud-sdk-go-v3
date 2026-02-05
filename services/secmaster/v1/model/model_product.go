package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Product struct {

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 资源规格编码
	ResourceSpecCode string `json:"resource_spec_code"`

	// 订阅数量
	ResourceSize int32 `json:"resource_size"`
}

func (o Product) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Product struct{}"
	}

	return strings.Join([]string{"Product", string(data)}, " ")
}
