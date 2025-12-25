package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportAlertsResponse Response Object
type ImportAlertsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *ImportAlertResponseBodyData `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ImportAlertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAlertsResponse struct{}"
	}

	return strings.Join([]string{"ImportAlertsResponse", string(data)}, " ")
}
