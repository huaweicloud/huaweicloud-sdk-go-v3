package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数异步配置返回结构体。
type ListFunctionAsyncInvokeConfigResult struct {

	// 函数URN。
	FuncUrn string `json:"func_urn" xml:"func_urn"`

	// 消息最大存活时长，取值范围[60，86400]。单位：秒。
	MaxAsyncEventAgeInSeconds int32 `json:"max_async_event_age_in_seconds" xml:"max_async_event_age_in_seconds"`

	// 异步调用失败后的最大重试次数，默认值为3。取值范围[0，8]。
	MaxAsyncRetryAttempts int32 `json:"max_async_retry_attempts" xml:"max_async_retry_attempts"`

	DestinationConfig *FuncAsyncDestinationConfig `json:"destination_config" xml:"destination_config"`

	// 异步调用配置的创建时间。
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 异步调用配置的最后更改时间。
	LastModified *string `json:"last_modified,omitempty" xml:"last_modified"`

	// 开启异步调用状态持久化
	EnableAsyncStatusLog *bool `json:"enable_async_status_log,omitempty" xml:"enable_async_status_log"`
}

func (o ListFunctionAsyncInvokeConfigResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionAsyncInvokeConfigResult struct{}"
	}

	return strings.Join([]string{"ListFunctionAsyncInvokeConfigResult", string(data)}, " ")
}
