package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertResponse Response Object
type ShowAlertResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *ShowAlertDetail `json:"data,omitempty"`

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
