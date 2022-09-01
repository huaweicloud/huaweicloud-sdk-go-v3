package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVersionAliasesRequest struct {

	// 函数的URN。
	FunctionUrn string `json:"function_urn" xml:"function_urn"`
}

func (o ListVersionAliasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionAliasesRequest struct{}"
	}

	return strings.Join([]string{"ListVersionAliasesRequest", string(data)}, " ")
}
