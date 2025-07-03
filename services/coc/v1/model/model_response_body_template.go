package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResponseBodyTemplate 通用返回体,用于文档生成。
type ResponseBodyTemplate struct {

	// 业务code，0 代表业务成功，其他数值代表有错误，详情请见错误码。
	Code *string `json:"code,omitempty"`

	// 服务编码。
	ProviderCode *string `json:"provider_code,omitempty"`

	// 错误信息，code为0，此值为空；code不为0，此处为具体的错误描述。
	Msg *string `json:"msg,omitempty"`
}

func (o ResponseBodyTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseBodyTemplate struct{}"
	}

	return strings.Join([]string{"ResponseBodyTemplate", string(data)}, " ")
}
