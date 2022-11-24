package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数异步配置返回结构体。
type FunctionAsyncConfig struct {

	// 消息最大存活时长，取值范围[60，86400]。单位：秒。
	MaxAsyncEventAgeInSeconds int32 `json:"max_async_event_age_in_seconds"`

	// 异步调用失败后的最大重试次数，默认值为3。取值范围[0，8]。
	MaxAsyncRetryAttempts int32 `json:"max_async_retry_attempts"`

	DestinationConfig *FuncAsyncDestinationConfig `json:"destination_config"`

	// 异步调用配置的创建时间。
	CreatedTime *string `json:"created_time,omitempty"`

	// 异步调用配置的最后更改时间。
	LastModified *string `json:"last_modified,omitempty"`
}

func (o FunctionAsyncConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionAsyncConfig struct{}"
	}

	return strings.Join([]string{"FunctionAsyncConfig", string(data)}, " ")
}
