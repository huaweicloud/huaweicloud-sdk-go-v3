package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyPlaybookVersionResponse Response Object
type CopyPlaybookVersionResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *PlaybookVersionInfo `json:"data,omitempty"`

	XRequestId *string `json:"X-request-id,omitempty"`

	XAPIDeprecationInfo *string `json:"X-API-Deprecation-Info,omitempty"`

	XAPIDeprecationDate *string `json:"X-API-Deprecation-Date,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CopyPlaybookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyPlaybookVersionResponse struct{}"
	}

	return strings.Join([]string{"CopyPlaybookVersionResponse", string(data)}, " ")
}
