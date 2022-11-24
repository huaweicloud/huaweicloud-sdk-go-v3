package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFunctionCodeRequest struct {

	// 函数的URN，详细解释见表1 FunctionGraph函数字段说明表的描述。
	FunctionUrn string `json:"function_urn"`

	Body *UpdateFunctionCodeRequestBody `json:"body,omitempty"`
}

func (o UpdateFunctionCodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFunctionCodeRequest struct{}"
	}

	return strings.Join([]string{"UpdateFunctionCodeRequest", string(data)}, " ")
}
