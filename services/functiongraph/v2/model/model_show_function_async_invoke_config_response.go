package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowFunctionAsyncInvokeConfigResponse struct {
	// 函数URN。

	FuncUrn *string `json:"func_urn,omitempty"`
	// 消息最大存活时长，取值范围[60，86400]。单位：秒。

	MaxAsyncEventAgeInSeconds *int32 `json:"max_async_event_age_in_seconds,omitempty"`
	// 异步调用失败后的最大重试次数，默认值为3。取值范围[0，8]。

	MaxAsyncRetryAttempts *int32 `json:"max_async_retry_attempts,omitempty"`

	DestinationConfig *FuncAsyncDestinationConfig `json:"destination_config,omitempty"`
	// 异步调用配置的创建时间。

	CreatedTime *string `json:"created_time,omitempty"`
	// 异步调用配置的最后更改时间。

	LastModified   *string `json:"last_modified,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFunctionAsyncInvokeConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionAsyncInvokeConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowFunctionAsyncInvokeConfigResponse", string(data)}, " ")
}
