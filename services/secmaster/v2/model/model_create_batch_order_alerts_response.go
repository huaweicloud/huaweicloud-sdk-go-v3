package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBatchOrderAlertsResponse Response Object
type CreateBatchOrderAlertsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *BatchOperateAlertResult `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBatchOrderAlertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchOrderAlertsResponse struct{}"
	}

	return strings.Join([]string{"CreateBatchOrderAlertsResponse", string(data)}, " ")
}
