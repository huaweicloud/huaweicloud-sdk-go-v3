package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVersionAliasesRequest Request Object
type ListVersionAliasesRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ListVersionAliasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionAliasesRequest struct{}"
	}

	return strings.Join([]string{"ListVersionAliasesRequest", string(data)}, " ")
}
