package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLtsLogDetailsRequest Request Object
type ShowLtsLogDetailsRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowLtsLogDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLtsLogDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowLtsLogDetailsRequest", string(data)}, " ")
}
