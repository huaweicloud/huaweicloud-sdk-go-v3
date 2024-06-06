package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRecyclePolicyRequest Request Object
type SetRecyclePolicyRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *SetRecyclePolicyRequestBody `json:"body,omitempty"`
}

func (o SetRecyclePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"SetRecyclePolicyRequest", string(data)}, " ")
}
