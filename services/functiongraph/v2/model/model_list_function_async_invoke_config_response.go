package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFunctionAsyncInvokeConfigResponse struct {

	// 函数异步配置列表。
	AsyncInvokeConfigs *[]ListFunctionAsyncInvokeConfigResult `json:"async_invoke_configs,omitempty" xml:"async_invoke_configs"`

	// 列表总数。
	Count *int64 `json:"count,omitempty" xml:"count"`

	PageInfo       *PageInfo `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int       `json:"-"`
}

func (o ListFunctionAsyncInvokeConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionAsyncInvokeConfigResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionAsyncInvokeConfigResponse", string(data)}, " ")
}
