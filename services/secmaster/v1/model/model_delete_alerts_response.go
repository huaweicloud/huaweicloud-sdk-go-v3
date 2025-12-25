package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertsResponse Response Object
type DeleteAlertsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *BatchOperateAlertResult `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o DeleteAlertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertsResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlertsResponse", string(data)}, " ")
}
