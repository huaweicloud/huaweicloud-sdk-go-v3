package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertResponse Response Object
type ShowAlertResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *AlertDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAlertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertResponse struct{}"
	}

	return strings.Join([]string{"ShowAlertResponse", string(data)}, " ")
}
