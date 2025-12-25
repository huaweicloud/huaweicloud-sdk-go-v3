package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePlaybookActionResponse Response Object
type UpdatePlaybookActionResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *ActionInfo `json:"data,omitempty"`

	XRequestId *string `json:"X-request-id,omitempty"`

	XAPIDeprecationInfo *string `json:"X-API-Deprecation-Info,omitempty"`

	XAPIDeprecationDate *string `json:"X-API-Deprecation-Date,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdatePlaybookActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePlaybookActionResponse struct{}"
	}

	return strings.Join([]string{"UpdatePlaybookActionResponse", string(data)}, " ")
}
