package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRemoveHeadersConfig 参数解释：要移除的请求头参数列表。
type UpdateRemoveHeadersConfig struct {

	// 参数解释：要移除的请求头参数列表。
	Configs []UpdateRemoveHeaderConfig `json:"configs"`
}

func (o UpdateRemoveHeadersConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRemoveHeadersConfig struct{}"
	}

	return strings.Join([]string{"UpdateRemoveHeadersConfig", string(data)}, " ")
}
