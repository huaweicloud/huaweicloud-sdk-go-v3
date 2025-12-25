package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookVersionResponse Response Object
type DeletePlaybookVersionResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePlaybookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookVersionResponse struct{}"
	}

	return strings.Join([]string{"DeletePlaybookVersionResponse", string(data)}, " ")
}
