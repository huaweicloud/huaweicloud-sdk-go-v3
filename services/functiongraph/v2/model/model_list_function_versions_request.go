package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionVersionsRequest Request Object
type ListFunctionVersionsRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 上一次查询到的最后的记录位置。
	Marker *string `json:"marker,omitempty"`

	// 每次查询获取的最大函数记录数量。
	Maxitems *string `json:"maxitems,omitempty"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListFunctionVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListFunctionVersionsRequest", string(data)}, " ")
}
