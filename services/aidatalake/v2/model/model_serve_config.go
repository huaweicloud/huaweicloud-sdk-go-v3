package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServeConfig **参数解释**：Ray Serve 应用配置。
type ServeConfig struct {

	// **参数解释**：Serve 应用列表。 **约束限制**：数量不超过3个，且各应用 name 与 route_prefix 均不可重复。
	Applications []ServeApplication `json:"applications"`
}

func (o ServeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServeConfig struct{}"
	}

	return strings.Join([]string{"ServeConfig", string(data)}, " ")
}
