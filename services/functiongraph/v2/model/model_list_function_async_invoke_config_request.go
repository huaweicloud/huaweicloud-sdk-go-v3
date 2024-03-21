package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionAsyncInvokeConfigRequest Request Object
type ListFunctionAsyncInvokeConfigRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 上一次查询到的最后的记录位置。
	Marker *string `json:"marker,omitempty"`

	// 每页显示的条目数量。 - 如果不提供该值或者提供的值等于0，则使用默认值：10，最大值100，大于100取值100。 - 如果该值小于0，则返回参数错误。
	Limit *string `json:"limit,omitempty"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListFunctionAsyncInvokeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionAsyncInvokeConfigRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionAsyncInvokeConfigRequest", string(data)}, " ")
}
