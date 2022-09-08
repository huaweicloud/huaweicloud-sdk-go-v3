package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// mock后端详情
type ApiMockCreate struct {

	// 描述信息。长度不超过255个字符 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 返回结果
	ResultContent *string `json:"result_content,omitempty"`

	// 版本。字符长度不超过64
	Version *string `json:"version,omitempty"`

	// 后端自定义认证ID
	AuthorizerId *string `json:"authorizer_id,omitempty"`
}

func (o ApiMockCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiMockCreate struct{}"
	}

	return strings.Join([]string{"ApiMockCreate", string(data)}, " ")
}
