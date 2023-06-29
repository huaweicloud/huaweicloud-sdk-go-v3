package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Retry 重试策略
type Retry struct {

	// 重试策略名称，在单个流程中，名称需要唯一
	Name string `json:"name"`

	// 重试间隔，单位：秒。若不传，默认为1
	Delay *int64 `json:"delay,omitempty"`

	// 最大重试次数，。若不传，默认为3
	MaxAttempts *int64 `json:"max_attempts,omitempty"`
}

func (o Retry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Retry struct{}"
	}

	return strings.Join([]string{"Retry", string(data)}, " ")
}
