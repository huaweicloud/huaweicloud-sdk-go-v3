package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicNetworkConfig **参数解释**：NoteBook网关类型 **约束限制**：不涉及。
type PublicNetworkConfig struct {

	// **参数解释**：NoteBook网络类型 **约束限制**： - SHARED：公共网络 - EXCLUSIVE：专属网络 - FORBIDDEN：禁用网络
	PublicNetworkType *string `json:"public_network_type,omitempty"`
}

func (o PublicNetworkConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicNetworkConfig struct{}"
	}

	return strings.Join([]string{"PublicNetworkConfig", string(data)}, " ")
}
