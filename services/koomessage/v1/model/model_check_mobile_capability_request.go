package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckMobileCapabilityRequest Request Object
type CheckMobileCapabilityRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *CheckMobileCapabilityRequestBody `json:"body,omitempty"`
}

func (o CheckMobileCapabilityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMobileCapabilityRequest struct{}"
	}

	return strings.Join([]string{"CheckMobileCapabilityRequest", string(data)}, " ")
}
