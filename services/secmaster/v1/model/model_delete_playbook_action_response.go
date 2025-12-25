package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookActionResponse Response Object
type DeletePlaybookActionResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`

	XRequestId *string `json:"X-request-id,omitempty"`

	XAPIDeprecationInfo *string `json:"X-API-Deprecation-Info,omitempty"`

	XAPIDeprecationDate *string `json:"X-API-Deprecation-Date,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o DeletePlaybookActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookActionResponse struct{}"
	}

	return strings.Join([]string{"DeletePlaybookActionResponse", string(data)}, " ")
}
