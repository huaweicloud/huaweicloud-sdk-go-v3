package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数异步配置请求体。
type UpdateFunctionAsyncInvokeConfigRequestBody struct {
	// 消息最大存活时长，取值范围[1，86400]，单位：秒，默认值为3600。

	MaxAsyncEventAgeInSeconds *int32 `json:"max_async_event_age_in_seconds,omitempty"`
	// 异步调用失败后的最大重试次数，默认值为1。取值范围[0，3]。

	MaxAsyncRetryAttempts *int32 `json:"max_async_retry_attempts,omitempty"`

	DestinationConfig *FuncAsyncDestinationConfig `json:"destination_config,omitempty"`
}

func (o UpdateFunctionAsyncInvokeConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionAsyncInvokeConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionAsyncInvokeConfigRequestBody", string(data)}, " ")
}
