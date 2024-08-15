package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRemoveHeadersConfig 参数解释：要移除的请求头参数列表。
type CreateRemoveHeadersConfig struct {

	// 参数解释：要移除的请求头参数列表。
	Configs []CreateRemoveHeaderConfig `json:"configs"`
}

func (o CreateRemoveHeadersConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRemoveHeadersConfig struct{}"
	}

	return strings.Join([]string{"CreateRemoveHeadersConfig", string(data)}, " ")
}
