package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFunctionRequest Request Object
type DeleteFunctionRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。  不允许删除函数的latest版本，如要删除整个函数（包含所有版本），提供不带任何版本号/别名的urn，如： urn:fss:xxxxxxxx:7aad83af3e8d42e99ac194e8419e2c9b:function:default:test
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o DeleteFunctionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionRequest struct{}"
	}

	return strings.Join([]string{"DeleteFunctionRequest", string(data)}, " ")
}
