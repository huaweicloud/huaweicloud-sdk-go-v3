package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAlertsResponse Response Object
type ChangeAlertsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *AlertDetail `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ChangeAlertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAlertsResponse struct{}"
	}

	return strings.Join([]string{"ChangeAlertsResponse", string(data)}, " ")
}
