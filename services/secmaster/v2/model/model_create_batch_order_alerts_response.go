package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBatchOrderAlertsResponse Response Object
type CreateBatchOrderAlertsResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *BatchOrderAlertResult `json:"data,omitempty"`

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
