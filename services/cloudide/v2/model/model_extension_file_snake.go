package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionFileSnake struct {

	// 资源类型
	AssetType *string `json:"asset_type,omitempty"`

	// 资源地址
	Source *string `json:"source,omitempty"`
}

func (o ExtensionFileSnake) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionFileSnake struct{}"
	}

	return strings.Join([]string{"ExtensionFileSnake", string(data)}, " ")
}
