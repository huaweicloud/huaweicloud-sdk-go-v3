package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveHeadersConfig 参数解释：要移除的请求头参数列表。
type RemoveHeadersConfig struct {

	// 参数解释：要移除的请求头参数列表。
	Configs []RemoveHeaderConfig `json:"configs"`
}

func (o RemoveHeadersConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveHeadersConfig struct{}"
	}

	return strings.Join([]string{"RemoveHeadersConfig", string(data)}, " ")
}
