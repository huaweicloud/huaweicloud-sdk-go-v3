package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIndicatorResponse Response Object
type DeleteIndicatorResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *IndicatorBatchOperateResponse `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteIndicatorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIndicatorResponse struct{}"
	}

	return strings.Join([]string{"DeleteIndicatorResponse", string(data)}, " ")
}
