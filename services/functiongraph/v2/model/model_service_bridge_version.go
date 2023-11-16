package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceBridgeVersion struct {

	// 代码包名
	Name *string `json:"name,omitempty"`

	// 代码版本
	Version *string `json:"version,omitempty"`

	// 代码所在obs路径
	CodeUrl *string `json:"code_url,omitempty"`
}

func (o ServiceBridgeVersion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceBridgeVersion struct{}"
	}

	return strings.Join([]string{"ServiceBridgeVersion", string(data)}, " ")
}
