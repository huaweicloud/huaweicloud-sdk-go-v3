package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResponseConfigDto 联邦配置
type ResponseConfigDto struct {

	// 额外添加的属性映射配置
	Properties map[string]ResponseSourceDetailsDto `json:"properties,omitempty"`

	Subject *ResponseSourceDetailsDto `json:"subject,omitempty"`

	// 中继状态
	RelayState *string `json:"relay_state,omitempty"`

	// 会话过期时间
	Ttl string `json:"ttl"`
}

func (o ResponseConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseConfigDto struct{}"
	}

	return strings.Join([]string{"ResponseConfigDto", string(data)}, " ")
}
