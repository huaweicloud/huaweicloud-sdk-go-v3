package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLayoutFieldResponse Response Object
type DeleteLayoutFieldResponse struct {

	// 响应码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLayoutFieldResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLayoutFieldResponse struct{}"
	}

	return strings.Join([]string{"DeleteLayoutFieldResponse", string(data)}, " ")
}
