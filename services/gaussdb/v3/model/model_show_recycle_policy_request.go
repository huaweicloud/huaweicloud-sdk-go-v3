package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecyclePolicyRequest Request Object
type ShowRecyclePolicyRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowRecyclePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowRecyclePolicyRequest", string(data)}, " ")
}
