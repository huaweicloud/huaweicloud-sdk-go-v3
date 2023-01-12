package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ChangeIncidentResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *IncidentDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeIncidentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIncidentResponse struct{}"
	}

	return strings.Join([]string{"ChangeIncidentResponse", string(data)}, " ")
}
