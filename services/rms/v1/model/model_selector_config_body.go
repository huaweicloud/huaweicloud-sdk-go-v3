package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// tracker选择器
type SelectorConfigBody struct {
	// 是否选择所有支持的资源

	AllSupported bool `json:"all_supported"`
	// 资源类型列表

	ResourceTypes []string `json:"resource_types"`
}

func (o SelectorConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectorConfigBody struct{}"
	}

	return strings.Join([]string{"SelectorConfigBody", string(data)}, " ")
}
