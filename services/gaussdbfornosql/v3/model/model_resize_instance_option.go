package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeInstanceOption struct {
	// 变更至新规格的资源规格编码。获取方法请参见查询所有实例规格信息中响应参数“flavors.spec_code”的值。

	TargetSpecCode string `json:"target_spec_code"`
}

func (o ResizeInstanceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceOption struct{}"
	}

	return strings.Join([]string{"ResizeInstanceOption", string(data)}, " ")
}
