package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionAsyncInvokeConfigResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionAsyncInvokeConfigResponse", string(data)}, " ")
}
