package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertResponse Response Object
type CreateAlertResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *AlertDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertResponse struct{}"
	}

	return strings.Join([]string{"CreateAlertResponse", string(data)}, " ")
}
