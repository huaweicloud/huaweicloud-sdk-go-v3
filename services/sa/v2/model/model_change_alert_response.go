package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAlertResponse Response Object
type ChangeAlertResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *AlertDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeAlertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAlertResponse struct{}"
	}

	return strings.Join([]string{"ChangeAlertResponse", string(data)}, " ")
}
