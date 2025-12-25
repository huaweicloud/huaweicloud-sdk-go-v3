package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookActionResponse Response Object
type CreatePlaybookActionResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 剧本Action列表
	Data *[]ActionInfo `json:"data,omitempty"`

	XRequestId *string `json:"X-request-id,omitempty"`

	XAPIDeprecationInfo *string `json:"X-API-Deprecation-Info,omitempty"`

	XAPIDeprecationDate *string `json:"X-API-Deprecation-Date,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreatePlaybookActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookActionResponse struct{}"
	}

	return strings.Join([]string{"CreatePlaybookActionResponse", string(data)}, " ")
}
