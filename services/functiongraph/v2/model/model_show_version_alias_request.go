package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVersionAliasRequest Request Object
type ShowVersionAliasRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 要查询的别名名称。
	AliasName string `json:"alias_name"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowVersionAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionAliasRequest struct{}"
	}

	return strings.Join([]string{"ShowVersionAliasRequest", string(data)}, " ")
}
