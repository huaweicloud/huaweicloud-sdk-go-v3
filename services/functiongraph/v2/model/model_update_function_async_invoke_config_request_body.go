package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数异步配置请求体。
type UpdateFunctionAsyncInvokeConfigRequestBody struct {

	// 消息最大存活时长，取值范围[1，86400]，单位：秒，默认值为3600。
	MaxAsyncEventAgeInSeconds *int32 `json:"max_async_event_age_in_seconds,omitempty" xml:"max_async_event_age_in_seconds"`

	// 异步调用失败后的最大重试次数，默认值为1。取值范围[0，3]。
	MaxAsyncRetryAttempts *int32 `json:"max_async_retry_attempts,omitempty" xml:"max_async_retry_attempts"`

	DestinationConfig *FuncAsyncDestinationConfig `json:"destination_config,omitempty" xml:"destination_config"`

	// 开启异步调用状态持久化
	EnableAsyncStatusLog *bool `json:"enable_async_status_log,omitempty" xml:"enable_async_status_log"`
}

func (o UpdateFunctionAsyncInvokeConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionAsyncInvokeConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionAsyncInvokeConfigRequestBody", string(data)}, " ")
}
