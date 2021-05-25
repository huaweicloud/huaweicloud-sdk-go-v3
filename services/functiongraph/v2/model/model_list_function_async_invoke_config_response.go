package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFunctionAsyncInvokeConfigResponse struct {
	// 函数异步配置列表。

	AsyncInvokeConfigs *[]ListFunctionAsyncInvokeConfigResult `json:"async_invoke_configs,omitempty"`
	// 列表总数。

	Count *int64 `json:"count,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListFunctionAsyncInvokeConfigResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFunctionAsyncInvokeConfigResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionAsyncInvokeConfigResponse", string(data)}, " ")
}
