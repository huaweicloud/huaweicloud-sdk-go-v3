package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAimPersonalTemplateResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 响应消息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAimPersonalTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAimPersonalTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteAimPersonalTemplateResponse", string(data)}, " ")
}
