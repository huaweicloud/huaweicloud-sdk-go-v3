package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFunctionUrlRequest Request Object
type UpdateFunctionUrlRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	Body *FunctionUrlRequestBody `json:"body,omitempty"`
}

func (o UpdateFunctionUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionUrlRequest struct{}"
	}

	return strings.Join([]string{"UpdateFunctionUrlRequest", string(data)}, " ")
}
