package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVmsTemplateStatusResponse Response Object
type ListVmsTemplateStatusResponse struct {

	// 状态码。
	Status *string `json:"status,omitempty"`

	// 响应信息。
	Message *string `json:"message,omitempty"`

	Data           *ListVmsTemplateStatusResponseMode `json:"data,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListVmsTemplateStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVmsTemplateStatusResponse struct{}"
	}

	return strings.Join([]string{"ListVmsTemplateStatusResponse", string(data)}, " ")
}
