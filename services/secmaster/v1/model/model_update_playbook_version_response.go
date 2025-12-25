package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePlaybookVersionResponse Response Object
type UpdatePlaybookVersionResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *PlaybookVersionInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePlaybookVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePlaybookVersionResponse struct{}"
	}

	return strings.Join([]string{"UpdatePlaybookVersionResponse", string(data)}, " ")
}
