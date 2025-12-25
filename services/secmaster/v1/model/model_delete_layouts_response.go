package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLayoutsResponse Response Object
type DeleteLayoutsResponse struct {

	// 响应码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLayoutsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLayoutsResponse struct{}"
	}

	return strings.Join([]string{"DeleteLayoutsResponse", string(data)}, " ")
}
