package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChatResourceConfig 资源配置。
type ChatResourceConfig struct {

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源数量
	CountResource *int32 `json:"count_resource,omitempty"`
}

func (o ChatResourceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatResourceConfig struct{}"
	}

	return strings.Join([]string{"ChatResourceConfig", string(data)}, " ")
}
