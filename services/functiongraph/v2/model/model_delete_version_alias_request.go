package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteVersionAliasRequest struct {
	// 函数的URN，详细解释见FunctionGraph函数模型的描述。

	FunctionUrn string `json:"function_urn"`
	// 要删除的别名名称。

	AliasName string `json:"alias_name"`
}

func (o DeleteVersionAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVersionAliasRequest struct{}"
	}

	return strings.Join([]string{"DeleteVersionAliasRequest", string(data)}, " ")
}
